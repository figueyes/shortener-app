package repository

import (
	"context"
	"fmt"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/persistence/mongo/config"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	ctx = context.TODO()
)

type MongoRepository struct {
	collectionName  string
	connection      *config.DbConnection
	mongoCollection *mongo.Collection
}

type OptionsRepository struct {
	Limit int64
	Skip  int64
}

func NewMongoRepository(collection string,
	connection *config.DbConnection) *MongoRepository {
	mongoRepository := &MongoRepository{
		collectionName: collection,
		connection:     connection,
	}
	client, err := mongoRepository.connection.GetConnection()
	if err != nil {
		panic(`cannot connect to mongo database`)
	}
	database := os.Getenv(fmt.Sprintf("%s%s", os.Getenv("APP_NAMESPACE"), "_MONGODB_DATABASE"))
	mongoRepository.mongoCollection = client.
		Database(database).
		Collection(mongoRepository.collectionName)
	return mongoRepository
}

func (b *MongoRepository) FindOne(query interface{}) (interface{}, error) {
	cursor := b.mongoCollection.FindOne(ctx, query)
	entity := make(map[string]interface{}, 0)
	err := cursor.Decode(&entity)
	if err != nil {
		if err.Error() != "" {
			return nil, nil
		}
		return nil, err
	}
	return entity, nil
}

func (b *MongoRepository) Save(body interface{}) (string, error) {
	result, err := b.mongoCollection.InsertOne(context.TODO(), body)
	if err != nil {
		log.Error("error trying to save data with: %s", err.Error())
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	log.Info("new object was created successfully with id: %s", id)
	return id, nil
}

func (b *MongoRepository) FindAndUpdate(query map[string]interface{}, body interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{}, 0)
	opt := options.FindOneAndUpdate().SetReturnDocument(1)
	err := b.mongoCollection.FindOneAndUpdate(ctx, query, body, opt).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Error("record does not exist: %s", err.Error())
		return nil, err
	} else if err != nil {
		log.Error("error trying to update into database with: %s", err.Error())
		return nil, err
	}
	log.Info("object was updated successfully: %s", utils.EntityToJson(query))
	return result, err
}
