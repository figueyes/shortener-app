package repositories

import "github.com/figueyes/shortener-app/app/shortener/domain/entities"

type ShortenerRepository interface {
	Create(short *entities.Short) (*entities.Short, error)
	FindByShortUrl(shortUrl string) (*entities.Short, error)
	Update(shortUrl string, short *entities.Short) (*entities.Short, error)
}
