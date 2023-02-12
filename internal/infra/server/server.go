package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
	db   *sql.DB
}

func New(e *echo.Echo, database *sql.DB) *Server {
	return &Server{
		echo: e,
		db:   database,
	}
}

func (s *Server) Start() error {
	s.echo.GET("/", hello)
	return nil
}
