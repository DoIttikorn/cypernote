package http

import (
	"github.com/Doittikorn/cypernote/internal/finance/usecase"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	financeUsecase usecase.FinanceUsecase
}

type FinanceHandler interface {
	Save()
}

func NewHttp(e *echo.Group, financeUsecase usecase.FinanceUsecase) {
	h := &Handler{
		financeUsecase: financeUsecase,
	}

	e.POST("/", h.Save)
}
