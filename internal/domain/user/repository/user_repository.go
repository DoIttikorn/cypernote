package repository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/domain/user"
)

type DB struct {
	db *sql.DB
}

func New(store *sql.DB) user.R {
	return &DB{
		db: store,
	}
}
