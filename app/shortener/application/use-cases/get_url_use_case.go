package use_cases

import (
	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/figueyes/shortener-app/app/shortener/domain/repositories"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/log"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
)

type getUrlUseCase struct {
	repository repositories.ShortenerRepository
}

func NewGetUrlUseCase(repository repositories.ShortenerRepository) *getUrlUseCase {
	return &getUrlUseCase{
		repository: repository,
	}
}
func (sh *getUrlUseCase) Get(short string) (*entities.Short, error) {
	shorted, err := sh.repository.FindByShortUrl(short)
	if err != nil {
		return nil, err
	}
	if !*shorted.IsEnable {
		log.Info("short is disable: %s", utils.EntityToJson(shorted))
		return nil, nil
	}
	return shorted, nil
}
