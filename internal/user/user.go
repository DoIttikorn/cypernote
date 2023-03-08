package user

import (
	"time"
)

type (
	// User is a struct that represents a user
	M struct {
		ID        float64   `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Age       int       `json:"age"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}

	// UserRepository is a repository interface for user
	R interface {
		GetByID(id float64) (float64, error)
	}

	// UserUsecase is a usecase interface for user
	U interface {
		GetByID(id float64) (float64, error)
	}
)
