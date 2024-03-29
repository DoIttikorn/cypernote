package route

import (
	"database/sql"
	"net/http"

	_financeRepository "github.com/Doittikorn/cypernote/internal/domain/finance/financeRepository"
	_financeUsecase "github.com/Doittikorn/cypernote/internal/domain/finance/financeUsecase"
	_financeHandler "github.com/Doittikorn/cypernote/internal/domain/finance/handler"
	_userRepository "github.com/Doittikorn/cypernote/internal/domain/user/userRepository"
	_userUsecase "github.com/Doittikorn/cypernote/internal/domain/user/userUsecase"
	"github.com/labstack/echo/v4"
)

type server struct {
	echo *echo.Echo
	db   *sql.DB
}

func New(router *echo.Echo, database *sql.DB) *server {
	return &server{
		echo: router,
		db:   database,
	}
}

func (s *server) RouteV1() {

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
	financeRoute := new(financeGroup, financeHandler)
	financeRoute.RouteFinance()

	// Tag group
	// tagGroup := v1.Group("/tag")
	// tagRepository := _tag.New(s.db)

}

func (s *server) Health() {
	s.echo.GET("/health", hello)
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}
