package https

import (
	"net/http"
	"time"

	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/pkg/errorc"
	"github.com/labstack/echo/v4"
)

func (h *FinanceHandler) save(c echo.Context) error {
	var m finance.M
	err := c.Bind(&m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorc.ErrBind)
	}

	m.Status = "Y"
	currentTime := time.Now()
	m.CreatedAt = currentTime
	m.UpdatedAt = currentTime

	err = h.usecase.Save(&m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorc.ResponseErr(err.Error()))
	}
	return c.JSON(http.StatusCreated, m)
}
