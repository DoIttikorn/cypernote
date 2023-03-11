package usecase

import (
	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/internal/user"
)

type usecase struct {
	FinanceRepository finance.R
	UserRepository    user.R
}

// initialize usecase for finance when call New Usecase can use
// finance repository and user repository
func New(financeRepo finance.R, userRepo user.R) finance.U {
	return &usecase{
		FinanceRepository: financeRepo,
		UserRepository:    userRepo,
	}
}

func (u *usecase) Update() {
}

func (u *usecase) Delete() {

}
