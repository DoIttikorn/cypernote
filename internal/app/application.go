package app

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/internal/domain/finance/financeRepository"
	"github.com/Doittikorn/cypernote/internal/domain/finance/financeUsecase"
	_financeHandler "github.com/Doittikorn/cypernote/internal/domain/finance/handler"
	"github.com/Doittikorn/cypernote/internal/domain/user/userRepository"
	"github.com/Doittikorn/cypernote/internal/domain/user/userUsecase"
)

type app struct {
	Handler handler
}

type handler struct {
	Finance finance.H
}

func InitializeApp(db *sql.DB) *app {
	return &app{
		Handler: handler{
			Finance: _financeHandler.New(
				financeUsecase.New(
					financeRepository.New(db),
					userUsecase.New(userRepository.New(db)),
				),
			),
		},
	}
}
