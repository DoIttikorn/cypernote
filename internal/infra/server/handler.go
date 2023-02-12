package server

import (
	"github.com/labstack/echo/v4"
)

// Handler
func hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
