package http

import (
	"github.com/Doittikorn/cypernote/internal/finance/usecase"
	"github.com/labstack/echo/v4"
)

type FinanceHandler struct {
	financeUsecase usecase.FinanceUsecase
}

type FinanceHttp interface {
	Save(c echo.Context) error
}

func NewHttp(e *echo.Group, financeUsecase usecase.FinanceUsecase) {
	h := &FinanceHandler{
		financeUsecase: financeUsecase,
	}

	e.POST("/", h.Save)
}
