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
	// filter for finance
	Filter struct {
		Type []string `json:"type"`
	}

	// repository interface for finance
	R interface {
		GetByUserID(userId int64, types []string) ([]M, error)
		Save(model *M) error
		Update()
		Delete()
	}

	// usecase interface for finance
	U interface {
		ExecuteGetByUserID(userID int64, filter *Filter) ([]M, error)
		Save(model *M) error
		Update()
		Delete()
	}

	// handler https interface for finance
	H interface {
		getByUserID(echo.Context) error
		delete(e echo.Context) error
		save(c echo.Context) error
	}
)
