package repository

import (
	"encoding/json"

	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/persistence/redis/config"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
	"github.com/go-redis/redis"
)

type RedisRepository struct {
	connection *config.DbConnection
	client     *redis.Client
}

func newRedisRepository(connection *config.DbConnection) *RedisRepository {
	redisRepository := &RedisRepository{
		connection: connection,
	}
	client, err := redisRepository.connection.GetConnection()
	if err != nil {
		panic(`cannot connect to redis database`)
	}
	redisRepository.client = client
	return redisRepository
}

func RedisRepositoryFactory() *RedisRepository {
	redisConfig := config.CreateDbRedisConnection()
	return newRedisRepository(redisConfig)
}

func (r *RedisRepository) Set(key, value string) error {
	err := r.client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepository) Get(key string) (*string, error) {
	value, err := r.client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (r *RedisRepository) Remove(key string) {
	_ = r.client.Del(key)
	log.Info("key was removed from cache: %s", key)
}

func (r *RedisRepository) DeleteFromCache(topic string) {
	subs := r.client.Subscribe(topic)
	payload := new(entities.Short)
	for {
		msg, err := subs.ReceiveMessage()
		if err != nil {
			log.Error("error trying to read message from queue")
			continue
		}
		log.Info("reading: %s", msg.String())
		utils.JsonToEntity(msg.Payload, &payload)
		r.Remove(payload.ShortUrl)
	}
}

func (r *RedisRepository) Publish(topic string, msg *entities.Short) error {
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Error("error trying to build payload: %s", err.Error())
		return err
	}
	p := r.client.Publish(topic, payload)
	if p.Err() != nil {
		log.Error("error trying to publish payload: %s", p.Err())
		return p.Err()
	}
	log.Info("message %s was sent to channel %s", utils.EntityToJson(msg), topic)
	return nil
}
