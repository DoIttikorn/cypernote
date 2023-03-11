package repository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/tag"
)

type Config struct {
	db *sql.DB
}

func New(db *sql.DB) tag.R {
	return &Config{
		db: db,
	}
}

func (c *Config) GetAll() ([]tag.M, error) {
	return nil, nil
}
func (c *Config) GetByID(id tag.TagID) (tag.M, error) {
	return tag.M{ID: 2131234, Name: "te"}, nil
}
func (c *Config) GetByName(name tag.TagName) (tag.M, error) {
	return tag.M{ID: 2131234, Name: "te"}, nil
}
func (c *Config) Create(tag tag.M) error {
	return nil
}
func (c *Config) Update(tag tag.M) error {
	return nil
}
func (c *Config) Delete(id tag.TagID) error {
	return nil
}
