package handler

import (
	"net/http"
	"strconv"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/pkg/errorc"
	"github.com/labstack/echo/v4"
)

type financeHandler struct {
	usecase finance.U
}

func New(finance finance.U) finance.H {
	return &financeHandler{
		usecase: finance,
	}
}

func (h *financeHandler) Save(c echo.Context) error {
	var m finance.FinanceRequest
	err := c.Bind(&m)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorc.ErrBind)
	}

	finance, err := h.usecase.ExecuteSave(&m)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorc.ResponseErr(err.Error()))
	}
	return c.JSON(http.StatusCreated, finance)
}

func (h *financeHandler) GetByUserID(c echo.Context) error {
	id := c.Param("id")
	var filter = new(finance.Filter)

	if err := c.Bind(filter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if id == "" {
		return echo.NewHTTPError(http.StatusNotFound, errorc.ErrNotFound)
	}

	userID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errorc.ErrInvalidID)
	}
	finance, err := h.usecase.ExecuteGetByUserID(userID, filter)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errorc.ResponseErr(err.Error()))
	}

	c.JSON(http.StatusOK, finance)
	return nil
}

func (h *financeHandler) Delete(e echo.Context) error {
	return nil
}
