package http

import (
	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/internal/finance/usecase"
	"github.com/labstack/echo/v4"
)

type FinanceHandler struct {
	financeUsecase usecase.U
	filter         finance.Filter
}

func NewHttp(e *echo.Group, financeUsecase usecase.U) {
	h := &FinanceHandler{
		financeUsecase: financeUsecase,
		filter:         finance.Filter{},
	}

	e.POST("/", h.Save)

	userFinance := e.Group("/user/:id")
	userFinance.GET("", h.GetByUserID)
}
