package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	t.Parallel()
	tests := []struct {
		scenario string
		function func(*testing.T)
	}{
		{
			scenario: "decode large url successfully",
			function: testShortUrlSuccess,
		},
	}
	for _, test := range tests {
		t.Run(test.scenario, test.function)
	}
}

func testShortUrlSuccess(t *testing.T) {
	length := 7
	testing1 := "https://www.test.com/url_with_many_characters/test/0123456789"
	shortOne := ShortInput(testing1, length)

	testing2 := "https://www.test.com/!@#$%^&*()_+"
	shortTwo := ShortInput(testing2, length)

	testing3 := "https://www.test.com"
	shortThree := ShortInput(testing3, length)

	testing4 := ""
	shortFour := ShortInput(testing4, length)

	testing5 := " 1"
	shortFive := ShortInput(testing5, length)

	testing6 := "1 "
	shortSix := ShortInput(testing6, length)

	assert.Equal(t, "bQZZxVD", shortOne)
	assert.Equal(t, "g7gcfPb", shortTwo)
	assert.Equal(t, "QmYRLJf", shortThree)
	assert.Equal(t, "UXg88A7", shortFour)
	assert.Equal(t, "6hXWBsP", shortFive)
	assert.Equal(t, "PjtBdG8", shortSix)
}
