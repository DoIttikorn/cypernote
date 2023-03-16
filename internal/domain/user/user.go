package user

import (
	"time"
)

type (
	// User is a struct that Model a user
	M struct {
		ID        int64     `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Age       int       `json:"age"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}
	// ID is a type that represents a user id
	UserID = int64

	// UserRepository is a repository interface for user
	R interface {
		AuditUserByID(model UserID) error
		GetByID(model *M) (M, error)
	}

	// UserUsecase is a usecase interface for user
	U interface {
		GetByUserID(id UserID) (M, error)
		AuditUserByID(id UserID) error
	}
)
