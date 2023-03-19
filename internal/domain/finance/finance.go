package finance

import (
	"time"

	"github.com/labstack/echo/v4"
)

type (
	// model for finance for database
	M struct {
		ID         int64     `json:"id"`
		UserID     int64     `json:"user_id"`
		Amount     float32   `json:"amount"`
		Note       string    `json:"note"`
		Type       string    `json:"type"`
		Status     string    `json:"status" default:"Y"`
		DateTimeAt time.Time `json:"datetime_at"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
	// repository interface for finance
	R interface {
		GetByUserID(userId int64, types []string) ([]FinanceResponse, error)
		Save(model *M) error
		Update()
		Delete()
	}
	// filter for finance
	Filter struct {
		Type []string `json:"type"`
	}

	// handler https interface for finance
	H interface {
		GetByUserID(echo.Context) error
		Delete(e echo.Context) error
		Save(c echo.Context) error
	}
)

// about usecase
type (
	FinanceRequest struct {
		UserID     int64     `json:"user_id"`
		Amount     float32   `json:"amount"`
		Note       string    `json:"note"`
		Type       string    `json:"type"`
		DateTimeAt time.Time `json:"datetime_at"`
	}

	FinanceResponse struct {
		ID         int64     `json:"id"`
		Amount     float32   `json:"amount"`
		Note       string    `json:"note"`
		Type       string    `json:"type"`
		Status     string    `json:"status" default:"Y"`
		DateTimeAt time.Time `json:"datetime_at"`
	}
	// usecase interface for finance
	U interface {
		ExecuteGetByUserID(userID int64, filter *Filter) ([]FinanceResponse, error)
		ExecuteSave(model *FinanceRequest) (*FinanceResponse, error)
		ExecuteUpdate()
		ExecuteDelete()
	}
)
