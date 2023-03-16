package repository

import (
	"github.com/Doittikorn/cypernote/internal/domain/user"
)

func (c *DB) GetByID(model *user.M) (user.M, error) {
	// stmt, err := c.db.Prepare(`
	// 	SELECT id
	// 	FROM "user" u
	// 	WHERE u.id = $1
	// `)
	// if err != nil {
	// 	return 0, err
	// }
	// var userID float64
	// err = stmt.QueryRow(id).Scan(&userID)
	// if err != nil {
	// 	return 0, errors.New("user not found")
	// }

	return user.M{}, nil
}
