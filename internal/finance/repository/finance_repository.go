package repository

import (
	"database/sql"
	"log"

	"github.com/Doittikorn/cypernote/internal/finance"
)

type Config struct {
	db *sql.DB
}

type Repository interface {
	GetByUserID()
	Save(context *finance.M)
	Update()
}

func New(db *sql.DB) Repository {
	return &Config{
		db: db,
	}
}

func (c *Config) GetByUserID() {
}

func (c *Config) Save(m *finance.M) {
	// save to database postgres
	query, err := c.db.Prepare(`
	INSERT INTO finance (user_id, amount, note, type, description) 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`)
	if err != nil {
		log.Println(err)
	}

	execute, err := query.Query(m.UserID, m.Amount, m.Note, m.Type, m.CreatedAt)
	if err != nil {
		log.Println(err)
	}

	// scan to struct
	err = execute.Scan(&m.ID)
	if err != nil {
		log.Println(err)
	}
}

func (c *Config) Update() {
}
