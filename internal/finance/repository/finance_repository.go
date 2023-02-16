package repository

import (
	"database/sql"
	"log"

	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/pkg/errorc"
)

type Config struct {
	db *sql.DB
}

type R interface {
	GetByUserID(userId float64) ([]finance.M, error)
	Save(context *finance.M) error
	Update()
}

func New(db *sql.DB) R {
	return &Config{
		db: db,
	}
}

func (c *Config) GetByUserID(userId float64) ([]finance.M, error) {
	var finances []finance.M
	query, err := c.db.Prepare(`
		SELECT id, user_id, amount, note, type, status, datetime_at, created_at
		FROM finance
		WHERE user_id = $1
	`)
	if err != nil {
		log.Println(err)
		return finances, errorc.New("prepare query error")
	}
	rows, err := query.Query(userId)
	if err != nil {
		log.Println(err)
		return finances, errorc.New("query error")
	}

	defer rows.Close()

	for rows.Next() {
		var finance finance.M
		err = rows.Scan(&finance.ID, &finance.UserID, &finance.Amount, &finance.Note, &finance.Type, &finance.Status, &finance.DateTimeAt, &finance.CreatedAt)
		if err != nil {
			log.Println(err)
			return finances, errorc.New("scan error")
		}
		finances = append(finances, finance)
	}

	return finances, nil
}

func (c *Config) Save(m *finance.M) error {
	// save to database postgres
	query, err := c.db.Prepare(`
		INSERT INTO finance (user_id, amount, note, type, status, datetime_at, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`)
	if err != nil {
		log.Println(err)
		return errorc.New("prepare query error")
	}
	err = query.QueryRow(m.UserID, m.Amount, m.Note, m.Type, m.Status, m.DateTimeAt, m.CreatedAt).Scan(&m.ID)
	if err != nil {
		log.Println(err)
		return errorc.New("query error")
	}

	return nil
}

func (c *Config) Update() {
}
