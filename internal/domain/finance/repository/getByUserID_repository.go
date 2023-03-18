package repository

import (
	"errors"
	"log"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/lib/pq"
)

func (c *DB) GetByUserID(userId int64, types []string) ([]finance.M, error) {
	var finances []finance.M
	if len(types) == 0 {
		types = []string{"income", "expense"}
	}

	query, err := c.db.Prepare(`
		SELECT id, user_id, amount, note, type, status, datetime_at, created_at, updated_at
		FROM finance
		WHERE user_id = $1 AND type = ANY ($2)
	`)

	if err != nil {
		log.Println(err)
		return finances, errors.New("prepare query error")
	}
	rows, err := query.Query(userId, pq.Array(types))
	if err != nil {
		log.Println(err)
		return finances, errors.New("query error")
	}

	defer rows.Close()

	for rows.Next() {
		var finance finance.M
		err = rows.Scan(&finance.ID, &finance.UserID, &finance.Amount, &finance.Note, &finance.Type, &finance.Status, &finance.DateTimeAt, &finance.CreatedAt, &finance.UpdatedAt)
		if err != nil {
			log.Println(err)
			return finances, errors.New("scan error")
		}
		finances = append(finances, finance)
	}

	return finances, nil
}
