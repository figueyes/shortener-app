package use_cases

import (
	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/figueyes/shortener-app/app/shortener/domain/repositories"
)

type modifyShortUseCase struct {
	repository repositories.ShortenerRepository
}

func NewModifyShortUseCase(repository repositories.ShortenerRepository) *modifyShortUseCase {
	return &modifyShortUseCase{
		repository: repository,
	}
}

func (sh *modifyShortUseCase) Modify(shortUrl string, short *entities.Short) (*entities.Short, error) {
	updated, err := sh.repository.Update(shortUrl, short)
	if err != nil {
		return nil, err
	}
	return updated, nil
}
