package cache

import "github.com/figueyes/shortener-app/app/shortener/domain/entities"

type Cache interface {
	Set(key, value string) error
	Get(key string) (*string, error)
	Publish(topic string, msg *entities.Short) error
}
