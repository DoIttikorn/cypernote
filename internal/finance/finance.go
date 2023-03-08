package finance

import (
	"time"
)

type (
	// model for finance
	M struct {
		ID         float64   `json:"id"`
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
		GetByUserID(userId float64, types []string) ([]M, error)
		Save(context *M) error
		Update()
	}

	// usecase interface for finance
	U interface {
		GetByUserID(userID float64, filter *Filter) ([]M, error)
		Save(model *M) error
		Update()
	}
)
