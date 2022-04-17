package use_cases

import (
	"github.com/figueyes/shortener-app/app/shortener/application/use-cases/mocks"
	"github.com/figueyes/shortener-app/app/shortener/domain/entities"
	"github.com/figueyes/shortener-app/app/shortener/infrastructure/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestModifyShortUseCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		scenario string
		function func(*testing.T)
	}{
		{
			scenario: "should it disable url successfully",
			function: testModifyEnable,
		},
	}
	for _, test := range tests {
		t.Run(test.scenario, test.function)
	}
}

func testModifyEnable(t *testing.T) {
	mockRepo := new(mocks.RepositoryMock)
	mockRepo.On("Update", mock.AnythingOfType("string"), mock.AnythingOfType("*entities.Short")).
		Return(&entities.Short{
			ShortUrl:    "bQZZxVD",
			OriginalUrl: url,
			IsEnable:    utils.BoolAddr(false),
			User:        user,
		}, nil)
	useCase := NewModifyShortUseCase(mockRepo)
	updated, err := useCase.Modify("bQZZxVD", &entities.Short{IsEnable: utils.BoolAddr(false)})
	expected := &entities.Short{
		ShortUrl:    "bQZZxVD",
		OriginalUrl: url,
		IsEnable:    utils.BoolAddr(false),
		User:        user,
	}
	mockRepo.AssertNumberOfCalls(t, "Update", 1)
	assert.Equal(t, expected, updated)
	assert.Equal(t, "bQZZxVD", mockRepo.Calls[0].Arguments[0])
	assert.Equal(t, &entities.Short{IsEnable: utils.BoolAddr(false)}, mockRepo.Calls[0].Arguments[1])
	assert.NoError(t, err)

}
