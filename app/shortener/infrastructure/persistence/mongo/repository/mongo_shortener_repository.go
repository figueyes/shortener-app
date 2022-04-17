package repository

import (
	"errors"
	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/persistence/mongo/config"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/persistence/mongo/repository/model"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
)

type MongoShortenerRepository struct {
	repository *MongoRepository
}

func newMongoShortenerRepository(repository *MongoRepository) *MongoShortenerRepository {
	mongoShortenerRepository := &MongoShortenerRepository{repository: repository}
	return mongoShortenerRepository
}

func MongoRepositoryFactory(collection string) *MongoShortenerRepository {
	c := config.CreateDbMongoConnection()
	repository := NewMongoRepository(collection, c)
	return newMongoShortenerRepository(repository)
}

func (m *MongoShortenerRepository) Create(short *entities.Short) (*entities.Short, error) {
	mongoShort := &model.MongoShort{
		ShortUrl:    short.ShortUrl,
		OriginalUrl: short.OriginalUrl,
		IsEnable:    short.IsEnable,
		User:        short.User,
	}
	saved, err := m.repository.Save(mongoShort)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if len(saved) == 0 {
		log.Error("error trying save data in storage")
		return nil, errors.New("data storage exception")
	}
	log.Info("saved short: %s with id %s", short.OriginalUrl, short.ShortUrl)
	return short, nil
}

func (m *MongoShortenerRepository) FindByShortUrl(shortUrl string) (*entities.Short, error) {
	r, err := m.repository.FindOne(map[string]interface{}{
		"short_url": shortUrl,
	})
	if r == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	mongoShort := new(model.MongoShort)
	utils.ConvertEntity(r, &mongoShort)
	short := new(entities.Short)
	short.OriginalUrl = mongoShort.OriginalUrl
	short.ShortUrl = mongoShort.ShortUrl
	short.User = mongoShort.User
	short.IsEnable = mongoShort.IsEnable
	return short, nil
}

func (m *MongoShortenerRepository) Update(shortUrl string, short *entities.Short) (*entities.Short, error) {
	updatedShort := new(entities.Short)
	mongoShort := new(model.MongoShort)
	if short.ShortUrl != "" {
		mongoShort.ShortUrl = short.ShortUrl
	}
	if short.OriginalUrl != "" {
		mongoShort.OriginalUrl = short.OriginalUrl
	}
	if short.User != "" {
		mongoShort.User = short.User
	}
	if short.IsEnable != nil {
		mongoShort.IsEnable = short.IsEnable
	}
	updated, err := m.repository.FindAndUpdate(
		map[string]interface{}{"short_url": shortUrl},
		map[string]interface{}{"$set": mongoShort},
	)
	if err != nil {
		return nil, err
	}
	utils.ConvertEntity(updated, &updatedShort)
	return updatedShort, nil
}
