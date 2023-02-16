package http

import (
	"github.com/Doittikorn/cypernote/internal/finance/usecase"
	"github.com/labstack/echo/v4"
)

type FinanceHandler struct {
	financeUsecase usecase.U
}

type FinanceHttp interface {
	Save(c echo.Context) error
}

func NewHttp(e *echo.Group, financeUsecase usecase.U) {
	h := &FinanceHandler{
		financeUsecase: financeUsecase,
	}

	e.POST("/", h.Save)

	userFinance := e.Group("/user/:id")
	userFinance.GET("", h.GetByUserID)
}
