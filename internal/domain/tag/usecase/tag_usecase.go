package usecase

import (
	"github.com/Doittikorn/cypernote/internal/domain/tag"
)

type usecase struct {
	TagRepository tag.R
}

func New(tagRe tag.R) tag.U {
	return &usecase{
		TagRepository: tagRe,
	}
}

func (t *usecase) GetAll() ([]tag.M, error) {
	return nil, nil
}
func (t *usecase) GetByID(id tag.TagID) (tag.M, error) {
	return tag.M{}, nil
}
func (t *usecase) GetByName(name tag.TagName) (tag.M, error) {
	return tag.M{}, nil
}
func (t *usecase) Create(tag tag.M) error {
	return nil
}
func (t *usecase) Update(tag tag.M) error {
	return nil
}
func (t *usecase) Delete(id tag.TagID) error {
	return nil
}
