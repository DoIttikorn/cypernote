package repository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
)

// don't another package can access this struct
type financeRepositoryDB struct {
	db *sql.DB
}

// initialize repository for finance when call New Config can use
// database connection
// and can use function in R interface
// R is interface for finance repository
func New(store *sql.DB) *financeRepositoryDB {
	return &financeRepositoryDB{
		db: store,
	}
}

// financeRepositoryMock is financeRepositoryMock for finance repository
type financeRepositoryMock struct {
	finance []*finance.M
}

func NewMock() *financeRepositoryMock {
	return &financeRepositoryMock{
		finance: make([]*finance.M, 0),
	}
}
