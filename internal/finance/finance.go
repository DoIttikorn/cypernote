package finance

import "time"

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
)
