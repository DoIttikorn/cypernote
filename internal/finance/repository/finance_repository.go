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

type Repository interface {
	GetByUserID()
	Save(context *finance.M) error
	Update()
}

func New(db *sql.DB) Repository {
	return &Config{
		db: db,
	}
}

func (c *Config) GetByUserID() {
}

func (c *Config) Save(m *finance.M) error {
	// save to database postgres
	query, err := c.db.Prepare(`INSERT INTO finance (user_id, amount, note, type, status) 
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`)
	if err != nil {
		log.Println(err)
		return errorc.New("prepare query error")
	}

	err = query.QueryRow(m.UserID, m.Amount, m.Note, m.Type, m.Status).Scan(&m.ID)
	if err != nil {
		log.Println(err)
		return errorc.New("query error")
	}

	return nil
}

func (c *Config) Update() {
}
