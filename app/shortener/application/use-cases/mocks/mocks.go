package mocks

import (
	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) Create(short *entities.Short) (*entities.Short, error) {
	m := r.Called(short)
	if m.Error(1) != nil {
		return nil, m.Error(1)
	}
	return m.Get(0).(*entities.Short), nil
}

func (r *RepositoryMock) FindByShortUrl(shortUrl string) (*entities.Short, error) {
	m := r.Called(shortUrl)
	if m.Error(1) != nil {
		return nil, m.Error(1)
	}
	if m.Get(0) == nil {
		return nil, nil
	}
	return m.Get(0).(*entities.Short), nil
}

func (r *RepositoryMock) Update(shortUrl string, short *entities.Short) (*entities.Short, error) {
	m := r.Called(shortUrl, short)
	if m.Error(1) != nil {
		return nil, m.Error(1)
	}
	return m.Get(0).(*entities.Short), nil
}
