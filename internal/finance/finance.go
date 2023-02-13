package finance

type M struct {
	ID         float64 `json:"id"`
	UserID     int64   `json:"user_id"`
	Amount     float32 `json:"amount"`
	Note       string  `json:"note"`
	Type       string  `json:"type"`
	Status     string  `json:"status" default:"Y"`
	DateTimeAt string  `json:"datetime_at"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
