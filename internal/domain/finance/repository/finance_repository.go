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
func New(store *sql.DB) finance.R {
	return &DB{
		db: store,
	}
}
