package tagRepository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/domain/tag"
)

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) tag.R {
	return &DB{
		db: db,
	}
}

func (c *DB) GetAll() ([]tag.M, error) {
	return nil, nil
}
func (c *DB) GetByID(id tag.TagID) (tag.M, error) {
	return tag.M{ID: 2131234, Name: "te"}, nil
}
func (c *DB) GetByName(name tag.TagName) (tag.M, error) {
	return tag.M{ID: 2131234, Name: "te"}, nil
}
func (c *DB) Create(tag tag.M) error {
	return nil
}
func (c *DB) Update(tag tag.M) error {
	return nil
}
func (c *DB) Delete(id tag.TagID) error {
	return nil
}
