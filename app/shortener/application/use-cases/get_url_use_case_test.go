package use_cases

import (
	"github.com/figueyes/shortener-app/app/shortener/application/use-cases/mocks"
	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetUrlUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		scenario string
		function func(*testing.T)
	}{
		{
			scenario: "get long url successfully",
			function: testGetUrlSuccess,
		},
	}
	for _, test := range tests {
		t.Run(test.scenario, test.function)
	}
}

func testGetUrlSuccess(t *testing.T) {
	shortUrl := "bQZZxVD"
	mockRepo := new(mocks.RepositoryMock)
	mockRepo.On("FindByShortUrl", mock.AnythingOfType("string")).
		Return(&entities.Short{
			ShortUrl:    "bQZZxVD",
			OriginalUrl: url,
			IsEnable:    utils.BoolAddr(true),
			User:        user,
		}, nil)
	useCase := NewGetUrlUseCase(mockRepo)
	shorted, err := useCase.Get(shortUrl)
	expected := &entities.Short{
		ShortUrl:    "bQZZxVD",
		OriginalUrl: url,
		IsEnable:    utils.BoolAddr(true),
		User:        user,
	}
	mockRepo.AssertNumberOfCalls(t, "FindByShortUrl", 1)
	assert.Equal(t, *expected, *shorted)
	assert.Equal(t, shortUrl, mockRepo.Calls[0].Arguments[0])
	assert.NoError(t, err)

}
