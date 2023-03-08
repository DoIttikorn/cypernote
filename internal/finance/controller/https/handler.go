package https

import (
	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/labstack/echo/v4"
)

type FinanceHandler struct {
	financeUsecase finance.U
	filter         finance.Filter
}

func NewHttp(e *echo.Group, usecase finance.U) {
	h := &FinanceHandler{
		financeUsecase: usecase,
		filter:         finance.Filter{},
	}

	e.POST("/", h.Save)

	userFinance := e.Group("/user/:id")
	userFinance.GET("", h.GetByUserID)
}
