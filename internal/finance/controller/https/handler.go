package https

import (
	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/labstack/echo/v4"
)

type FinanceHandler struct {
	financeUsecase finance.U
	filter         finance.Filter
}

func NewHttp(e *echo.Group, financeUsecase finance.U) {
	h := &FinanceHandler{
		financeUsecase: financeUsecase,
		filter:         finance.Filter{},
	}

	e.POST("/", h.Save)

	userFinance := e.Group("/user/:id")
	userFinance.GET("", h.GetByUserID)
}
