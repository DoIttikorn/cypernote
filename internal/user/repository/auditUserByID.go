package repository

import (
	"errors"

	"github.com/Doittikorn/cypernote/internal/user"
)

func (c *Config) AuditUserByID(userId user.UserID) error {
	stmt, err := c.db.Prepare(`
		SELECT id
		FROM "user" u
		WHERE u.id = $1
	`)
	if err != nil {
		return err
	}

	var userID float64
	err = stmt.QueryRow(userId).Scan(&userID)
	if err != nil {
		return errors.New("user not found")
	}

	return nil
}
