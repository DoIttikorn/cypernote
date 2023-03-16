package tag

type (
	// Model is a struct that Model a tag
	M struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	TagName = string
	TagID   = int64

	// TagRepository is a repository interface for tag
	R interface {
		GetAll() ([]M, error)
		GetByID(id TagID) (M, error)
		GetByName(name TagName) (M, error)
		Create(tag M) error
		Update(tag M) error
		Delete(id TagID) error
	}

	// TagUsecase is a usecase interface for tag
	U interface {
		GetAll() ([]M, error)
		GetByID(id TagID) (M, error)
		GetByName(name TagName) (M, error)
		Create(tag M) error
		Update(tag M) error
		Delete(id TagID) error
	}
)
