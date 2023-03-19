package financeRepository

import (
	"errors"
	"time"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/lib/pq"
)

func (c *financeRepositoryDB) GetByUserID(userId int64, types []string) ([]finance.FinanceResponse, error) {
	var finances []finance.FinanceResponse
	if len(types) == 0 {
		types = []string{"income", "expense"}
	}

	query, err := c.db.Prepare(`
		SELECT id, amount, note, type, status, datetime_at
		FROM finance
		WHERE user_id = $1 AND type = ANY ($2)
	`)

	if err != nil {
		return finances, errors.New("prepare query error")
	}
	rows, err := query.Query(userId, pq.Array(types))
	if err != nil {
		return finances, errors.New("query error")
	}

	defer rows.Close()

	for rows.Next() {
		var finance finance.FinanceResponse
		err = rows.Scan(&finance.ID, &finance.Amount, &finance.Note, &finance.Type, &finance.Status, &finance.DateTimeAt)
		if err != nil {
			return finances, errors.New("error result data")
		}
		finances = append(finances, finance)
	}

	return finances, nil
}

func (m *financeRepositoryMock) GetByUserID(userId int64, types []string) ([]finance.FinanceResponse, error) {
	return []finance.FinanceResponse{
		{
			ID:         1,
			Amount:     100,
			Note:       "test",
			Type:       "income",
			Status:     "Y",
			DateTimeAt: time.Now(),
		},
	}, nil
}
