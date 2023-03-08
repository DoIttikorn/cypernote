package tag

type (
	// Tag is a struct that represents a tag
	Tag struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	// TagRepository is a repository interface for tag
	TagRepository interface {
		GetAll() ([]Tag, error)
		GetByID(id int64) (Tag, error)
		GetByName(name string) (Tag, error)
		Create(tag Tag) error
		Update(tag Tag) error
		Delete(id int64) error
	}

	// TagUsecase is a usecase interface for tag
	TagUsecase interface {
		GetAll() ([]Tag, error)
		GetByID(id int64) (Tag, error)
		GetByName(name string) (Tag, error)
		Create(tag Tag) error
		Update(tag Tag) error
		Delete(id int64) error
	}
)
