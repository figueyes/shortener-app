package use_cases

import "github.com/figueyes/shortener-app/app/shortener/domain/entities"

type Short interface {
	Short(short *entities.Short) (*entities.Short, error)
}

type Get interface {
	Get(short string) (*entities.Short, error)
}

type Modify interface {
	Modify(shortUrl string, short *entities.Short) (*entities.Short, error)
}
