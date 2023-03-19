package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/pkg/errorc"
	"github.com/labstack/echo/v4"
)

type financeHandler struct {
	usecase finance.U
}

func New(finance finance.U) *financeHandler {
	return &financeHandler{
		usecase: finance,
	}
}

func (h *financeHandler) Route(e *echo.Group) {
	// start endpoint
	// v1/finance
	e.POST("/", h.save)
	e.DELETE("/", h.delete)
	// e.POST("/", h.Update)

	// user group for finance
	userFinance := e.Group("/user/:id")

	userFinance.GET("", h.getByUserID)
}

func (h *financeHandler) save(c echo.Context) error {
	var m finance.M
	err := c.Bind(&m)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorc.ErrBind)
	}

	m.Status = "Y"
	currentTime := time.Now()
	m.CreatedAt = currentTime
	m.UpdatedAt = currentTime

	err = h.usecase.ExecuteSave(&m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorc.ResponseErr(err.Error()))
	}
	return c.JSON(http.StatusCreated, m)
}

func (h *financeHandler) getByUserID(c echo.Context) error {
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
	finance, err := h.usecase.ExecuteGetByUserID(userID, filter)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorc.ResponseErr(err.Error()))
	}

	c.JSON(http.StatusOK, finance)
	return nil
}

func (h *financeHandler) delete(e echo.Context) error {
	return nil
}
