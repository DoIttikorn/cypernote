package repository

import (
	"errors"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
)

func (c *DB) Save(m *finance.M) error {
	query, err := c.db.Prepare(`
		INSERT INTO finance (user_id, amount, note, type, status, datetime_at, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`)
	println(query)
	if err != nil {
		return errors.New("prepare query error")
	}
	err = query.QueryRow(m.UserID, m.Amount, m.Note, m.Type, m.Status, m.DateTimeAt, m.CreatedAt, m.UpdatedAt).Scan(&m.ID)
	if err != nil {
		return errors.New("query error")
	}

	return nil
}
