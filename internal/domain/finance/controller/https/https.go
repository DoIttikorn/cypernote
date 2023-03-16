package https

import (
	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/labstack/echo/v4"
)

type FinanceHandler struct {
	usecase finance.U
	filter  finance.Filter
}

func NewHttp(e *echo.Group, usecase finance.U) {

	h := &FinanceHandler{
		usecase: usecase,
		filter:  finance.Filter{},
	}

	e.POST("/", h.save)
	e.DELETE("/", h.delete)
	// e.POST("/", h.Update)

	// user group for finance
	userFinance := e.Group("/user/:id")

	userFinance.GET("", h.getByUserID)

}
