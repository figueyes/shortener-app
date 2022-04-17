package config

import (
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
	"github.com/go-redis/redis"
)

type Connection interface {
	GetConnection() (*redis.Client, error)
	CloseConnection()
}

type DbConnection struct {
	opts *RedisOptions
}

func NewRedisConnection(options ...*RedisOptions) *DbConnection {
	databaseOptions := MergeOptions(options...)
	return &DbConnection{
		opts: databaseOptions,
	}
}

func (r *DbConnection) GetConnection() (*redis.Client, error) {
	connection := redis.NewClient(&redis.Options{
		Addr:     *r.opts.Addr,
		Password: *r.opts.Password,
		DB:       *r.opts.DB,
	},
	)
	pong, err := connection.Ping().Result()
	if err != nil {
		return nil, err
	}
	log.Info("Redis connection test: %s!!!", pong)
	return connection, nil
}
