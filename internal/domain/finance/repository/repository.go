package repository

import (
	"database/sql"

	"github.com/Doittikorn/cypernote/internal/domain/finance"
)

// don't another package can access this struct
type financeDB struct {
	db *sql.DB
}

// initialize repository for finance when call New Config can use
// database connection
// and can use function in R interface
// R is interface for finance repository
func New(store *sql.DB) *financeDB {
	return &financeDB{
		db: store,
	}
}

// mock is mock for finance repository
type mock struct {
	finance []*finance.M
}

func NewMock() *mock {
	return &mock{
		finance: make([]*finance.M, 0),
	}
}
