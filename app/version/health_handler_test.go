package version

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler_HealthCheck(t *testing.T) {
	t.Parallel()
	t.Run("should it run health handler successfully", func(t *testing.T) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		recorder := httptest.NewRecorder()
		context := e.NewContext(request, recorder)
		version := "1.0.0"
		appName := "test"
		author := "test"
		health := NewHealthHandler(e, version, appName, author)
		err := health.HealthCheck(context)
		assert.NoError(t, err)
	})
}
