package repository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/finance"
)

type Config struct {
	db *sql.DB
}

// initialize repository for finance when call New Config can use
// database connection
// and can use function in R interface
// R is interface for finance repository
func New(db *sql.DB) finance.R {
	return &Config{
		db: db,
	}
}
func (c *Config) Delete() {
}

func (c *Config) Update() {
}
