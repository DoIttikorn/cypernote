package http

import (
	"net/http"

	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/pkg/errorc"
	"github.com/labstack/echo/v4"
)

func (h *FinanceHandler) Save(c echo.Context) error {
	var m finance.M
	err := c.Bind(&m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorc.NewErrorRes(http.StatusBadRequest, "Bad Request"))
	}

	m.Status = "Y"
	err = h.financeUsecase.Save(&m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorc.NewErrorRes(http.StatusInternalServerError, err.Error()))
	}
	return c.JSON(http.StatusCreated, m)
}
