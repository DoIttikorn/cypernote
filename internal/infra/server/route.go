package server

import (
	"database/sql"

	_financeHttp "github.com/Doittikorn/cypernote/internal/finance/controller/http"
	_financeRepository "github.com/Doittikorn/cypernote/internal/finance/repository"
	_financeUsecase "github.com/Doittikorn/cypernote/internal/finance/usecase"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
	db   *sql.DB
}

func New(router *echo.Echo, database *sql.DB) *Server {
	return &Server{
		echo: router,
		db:   database,
	}
}

func (s *Server) Route() {

	v1 := s.echo.Group("/v1")
	// Routes

	//* Finance group
	financeGroup := v1.Group("/finance")
	financeRepository := _financeRepository.New(s.db)
	financeUsecase := _financeUsecase.New(financeRepository)
	_financeHttp.NewHttp(financeGroup, financeUsecase)

	v1.GET("/", hello)

}

func hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
