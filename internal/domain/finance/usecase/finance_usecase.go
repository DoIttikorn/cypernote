package usecase

import (
	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/internal/domain/user"
)

type usecase struct {
	financeRepository finance.R
	userUsecase       user.U
}

// initialize usecase for finance when call New Usecase can use
// finance repository and user repository
func New(financeRepo finance.R, userRepo user.U) finance.U {
	return &usecase{
		financeRepository: financeRepo,
		userUsecase:       userRepo,
	}
}

func (u *usecase) Update() {
}

func (u *usecase) Delete() {

}
