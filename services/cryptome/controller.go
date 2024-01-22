package cryptome

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler
func (ctrl *Control) hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! cryptome")
}
