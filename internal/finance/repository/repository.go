package repository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/finance"
)

type Config struct {
	db *sql.DB
}

func New(db *sql.DB) finance.R {
	return &Config{
		db: db,
	}
}

func (c *Config) Update() {
}
