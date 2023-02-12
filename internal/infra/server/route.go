package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server *echo.Echo
	db     *sql.DB
}

func New(router *echo.Echo, database *sql.DB) *Server {
	return &Server{
		server: router,
		db:     database,
	}
}

func (s *Server) Route() {
	// Routes
	s.server.GET("/", hello)
}

func hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
