package use_cases

import (
	"github.com/figueyes/shortener-app/app/shortener/application/use-cases/mocks"
	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	url  = "https://www.test.com/url_with_many_characters/test/0123456789"
	user = "test-user"
)

func TestShortUrlUseCase(t *testing.T) {

	t.Parallel()
	tests := []struct {
		scenario string
		function func(*testing.T)
	}{
		{
			scenario: "short large url successfully",
			function: testShortUrlSuccessfully,
		},
	}
	for _, test := range tests {
		t.Run(test.scenario, test.function)
	}
}

func testShortUrlSuccessfully(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	mockRepo.On("FindByShortUrl", mock.AnythingOfType("string")).
		Return(nil, nil)
	mockRepo.On("Create", mock.AnythingOfType("*entities.Short")).
		Return(&entities.Short{
			ShortUrl:    "bQZZxVD",
			OriginalUrl: url,
			IsEnable:    utils.BoolAddr(true),
			User:        user,
		}, nil)
	useCase := NewShortUrlUseCase(mockRepo)
	short := &entities.Short{
		OriginalUrl: url,
		User:        user,
	}
	shorted, err := useCase.Short(short)
	expectedShort := &entities.Short{
		ShortUrl:    "bQZZxVD",
		OriginalUrl: url,
		IsEnable:    utils.BoolAddr(true),
		User:        user,
	}
	mockRepo.AssertNumberOfCalls(t, "Create", 1)
	assert.Equal(t, *expectedShort, *shorted)
	assert.Equal(t, short.ShortUrl, mockRepo.Calls[0].Arguments[0])
	assert.Equal(t, short, mockRepo.Calls[1].Arguments[0])
	assert.NoError(t, err)
}
