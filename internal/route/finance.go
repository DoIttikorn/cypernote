package route

import (
	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/labstack/echo/v4"
)

type financeHandler struct {
	app     *echo.Group
	handler finance.H
}

func new(app *echo.Group, handler finance.H) *financeHandler {
	return &financeHandler{
		app:     app,
		handler: handler,
	}
}

func (s *financeHandler) RouteFinance() { // start endpoint
	// v1/finance
	s.app.POST("/", s.handler.Save)
	s.app.DELETE("/", s.handler.Delete)
	// e.POST("/", h.Update)

	// user group for finance
	userFinance := s.app.Group("/user/:id")

	userFinance.GET("", s.handler.GetByUserID)

}
