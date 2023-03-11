package repository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/user"
)

type Config struct {
	db *sql.DB
}

func New(db *sql.DB) user.R {
	return &Config{
		db: db,
	}
}
