package config

import (
	"context"
	"errors"
	"sync"

	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connection      *mongo.Client
	connectionError error
	mongoOnce       sync.Once
)

type Connection interface {
	GetConnection() (*mongo.Client, error)
	CloseConnection()
}

type DbConnection struct {
	opts *Options
	url  string
}

func NewMongoConnection(options ...*Options) *DbConnection {
	databaseOptions := MergeOptions(options...)
	url := databaseOptions.GetUrlConnection()
	if url == "" {
		log.Fatal(errors.New("error creating connection, url cannot be empty").Error())
	}
	return &DbConnection{
		opts: databaseOptions,
		url:  url,
	}
}

func (r *DbConnection) GetConnection() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(r.url)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			connectionError = err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			connectionError = err
		}
		connection = client
	})
	return connection, connectionError
}
