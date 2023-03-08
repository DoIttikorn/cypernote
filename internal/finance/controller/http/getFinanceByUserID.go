package http

import (
	"net/http"
	"strconv"

	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/pkg/errorc"
	"github.com/labstack/echo/v4"
)

func (h *FinanceHandler) GetFinanceByUserID(c echo.Context) error {
	id := c.Param("id")
	var filter = new(finance.Filter)

	if err := c.Bind(filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if id == "" {
		return c.JSON(http.StatusBadRequest, errorc.NewErrorRes(http.StatusBadRequest, "not get id"))
	}

	userID, err := strconv.ParseFloat(id, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, errorc.NewErrorRes(http.StatusBadRequest, "pattern id is not correct"))
	}

	finance, err := h.financeUsecase.GetByUserID(userID, filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorc.NewErrorRes(http.StatusInternalServerError, err.Error()))
	}

	c.JSON(http.StatusOK, finance)
	return nil
}
