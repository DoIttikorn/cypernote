package finance

type M struct {
	ID        float64 `json:"id"`
	UserID    int64   `json:"user_id"`
	Amount    float32 `json:"amount"`
	Note      string  `json:"note"`
	Type      string  `json:"type"`
	Status    string  `json:"status" default:"Y"`
	UpdatedBy string  `json:"updated_by"`
	CreatedAt string  `json:"created_at"`
}
