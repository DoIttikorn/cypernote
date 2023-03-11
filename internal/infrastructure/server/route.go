package server

import (
	"database/sql"

	_financeHttp "github.com/Doittikorn/cypernote/internal/finance/controller/https"
	_financeRepository "github.com/Doittikorn/cypernote/internal/finance/repository"
	_financeUsecase "github.com/Doittikorn/cypernote/internal/finance/usecase"
	_userRepository "github.com/Doittikorn/cypernote/internal/user/repository"
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

	// Routes
	v1 := s.echo.Group("/v1")

	//* User group
	// userGroup := v1.Group("/user")
	userRepository := _userRepository.New(s.db)

	//* Finance group
	financeGroup := v1.Group("/finance")
	financeRepository := _financeRepository.New(s.db)
	financeUsecase := _financeUsecase.New(financeRepository, userRepository)
	_financeHttp.NewHttp(financeGroup, financeUsecase)

	// Tag group
	// tagGroup := v1.Group("/tag")
	// tagRepository := _tag.New(s.db)

	// health check
	v1.GET("", hello)

}

func hello(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
