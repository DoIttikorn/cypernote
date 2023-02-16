package repository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/pkg/errorc"
)

type Config struct {
	db *sql.DB
}

type R interface {
	GetByID(id float64) (float64, error)
}

func New(db *sql.DB) R {
	return &Config{
		db: db,
	}
}

func (c *Config) GetByID(id float64) (float64, error) {
	stmt, err := c.db.Prepare(`
		SELECT id
		FROM "user" u
		WHERE u.id = $1
	`)
	if err != nil {
		return 0, err
	}
	var userID float64
	err = stmt.QueryRow(id).Scan(&userID)
	if err != nil {
		return 0, errorc.New("user not found")
	}

	return userID, nil
}
