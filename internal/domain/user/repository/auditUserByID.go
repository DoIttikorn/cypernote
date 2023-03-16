package repository

import (
	"github.com/Doittikorn/cypernote/internal/domain/user"
	"github.com/Doittikorn/cypernote/pkg/errorc"
)

func (c *DB) AuditUserByID(userId user.UserID) error {
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
		return errorc.ErrNotFound
	}

	return nil
}
