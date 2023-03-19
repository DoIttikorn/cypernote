package app

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
	_financeHandler "github.com/Doittikorn/cypernote/internal/domain/finance/handler"
	_financeRepository "github.com/Doittikorn/cypernote/internal/domain/finance/repository"
	_financeUsecase "github.com/Doittikorn/cypernote/internal/domain/finance/usecase"
	_userRepository "github.com/Doittikorn/cypernote/internal/domain/user/userRepository"
	_userUsecase "github.com/Doittikorn/cypernote/internal/domain/user/userUsecase"
)

type app struct {
	Handler handler
}

type handler struct {
	Finance finance.H
}

func InitializeApp(db *sql.DB) app {
	return app{
		Handler: handler{
			Finance: _financeHandler.New(
				_financeUsecase.New(
					_financeRepository.New(db),
					_userUsecase.New(_userRepository.New(db)),
				),
			),
		},
	}
}
