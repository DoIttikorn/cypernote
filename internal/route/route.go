package route

import (
	"database/sql"
	"net/http"

	_financeHandler "github.com/Doittikorn/cypernote/internal/domain/finance/handler"
	_financeRepository "github.com/Doittikorn/cypernote/internal/domain/finance/repository"
	_financeUsecase "github.com/Doittikorn/cypernote/internal/domain/finance/usecase"
	_userRepository "github.com/Doittikorn/cypernote/internal/domain/user/repository"
	_userUsecase "github.com/Doittikorn/cypernote/internal/domain/user/usecase"
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

func (s *Server) RouteV1() {

	// Routes
	v1 := s.echo.Group("/v1")

	//* User group
	// userGroup := v1.Group("/user")
	userRepository := _userRepository.New(s.db)
	userUsecase := _userUsecase.New(userRepository)

	//* Finance group
	financeGroup := v1.Group("/finance")
	financeRepository := _financeRepository.New(s.db)
	financeUsecase := _financeUsecase.New(financeRepository, userUsecase)
	financeHandler := _financeHandler.New(financeUsecase)
	financeHandler.Route(financeGroup)
	// route

	// Tag group
	// tagGroup := v1.Group("/tag")
	// tagRepository := _tag.New(s.db)

}

func (s *Server) Health() {
	s.echo.GET("/health", hello)
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
