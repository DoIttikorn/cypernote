package https

import (
	"net/http"
	"strconv"

	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/pkg/errorc"
	"github.com/labstack/echo/v4"
)

func (h *FinanceHandler) getByUserID(c echo.Context) error {
	id := c.Param("id")
	var filter = new(finance.Filter)

	if err := c.Bind(filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if id == "" {
		return c.JSON(http.StatusNotFound, errorc.ErrNotFound)
	}

	userID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, errorc.ErrInvalidID)
	}
	finance, err := h.usecase.GetByUserID(userID, filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorc.ResponseErr(err.Error()))
	}

	c.JSON(http.StatusOK, finance)
	return nil
}
