package use_cases

import (
	"fmt"
	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/figueyes/shortener-app/app/shortener/domain/exceptions"
	"github.com/figueyes/shortener-app/app/shortener/domain/repositories"
	"github.com/figueyes/shortener-app/app/shortener/domain/services"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
	"os"
	"strconv"
)

type shortUrlUseCase struct {
	length     int
	repository repositories.ShortenerRepository
}

func NewShortUrlUseCase(repository repositories.ShortenerRepository) *shortUrlUseCase {
	lenInt, err := strconv.Atoi(os.Getenv("LENGTH_SHORT_URL"))
	if err != nil || lenInt == 0 {
		lenInt = 7
	}
	return &shortUrlUseCase{
		length:     lenInt,
		repository: repository,
	}
}

func (sh *shortUrlUseCase) Short(short *entities.Short) (*entities.Short, error) {
	if short == nil || len(short.OriginalUrl) == 0 {
		return nil, exceptions.InvalidShortToCreateException
	}

	shortUrl := services.ShortInput(fmt.Sprintf("%s%s", short.OriginalUrl, short.User), sh.length)

	alreadyCreated, err := sh.repository.FindByShortUrl(shortUrl)
	if err != nil {
		return nil, exceptions.DataBaseException
	}
	if alreadyCreated != nil {
		return alreadyCreated, nil
	}

	(*short).ShortUrl = shortUrl
	(*short).IsEnable = utils.BoolAddr(true)

	created, err := sh.repository.Create(short)
	if err != nil {
		return nil, exceptions.DataBaseException
	}
	return created, nil
}
