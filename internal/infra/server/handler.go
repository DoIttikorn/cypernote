package server

import (
	"github.com/labstack/echo/v4"
)

func New(e *echo.Echo) *Server {
	return &Server{
		echo: e,
	}
}

func (s *Server) Start() error {
	s.echo.GET("/", hello)
	return nil
}

// Handler
func hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
