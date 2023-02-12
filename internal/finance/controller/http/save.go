package http

import "github.com/labstack/echo/v4"

func (h *Handler) Save(c echo.Context) error {
	return c.JSON(200, "Hello, World!")
}
