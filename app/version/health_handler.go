package version

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
}

type healthHandler struct {
	version string
	appName string
	author  string
}

func NewHealthHandler(e *echo.Echo,
	version string,
	appName string,
	author string,
) *healthHandler {
	h := &healthHandler{
		version: version,
		appName: appName,
		author:  author,
	}
	e.GET(fmt.Sprintf("%s/version", os.Getenv("BASE_PATH")), h.HealthCheck)
	return h
}

func (h *healthHandler) HealthCheck(c echo.Context) error {

	healthCheck := HealthCheck{
		App:     h.appName,
		Version: h.version,
		Env:     os.Getenv("ENVIRONMENT_TYPE"),
		Author:  h.author,
	}
	return c.JSON(http.StatusOK, healthCheck)
}
