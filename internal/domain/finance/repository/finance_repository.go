package repository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
)

type DB struct {
	db *sql.DB
}

// initialize repository for finance when call New Config can use
// database connection
// and can use function in R interface
// R is interface for finance repository
func New(db *sql.DB) finance.R {
	return &DB{
		db: db,
	}
}
func (c *DB) Delete() {
}

func (c *DB) Update() {
}
