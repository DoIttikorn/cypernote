package usecase

import (
	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/internal/domain/user"
)

// don't another package can access this struct
type financeUsecase struct {
	financeRepository finance.R
	userUsecase       user.U
}

// initialize usecase for finance when call New Usecase can use
// finance repository and user repository
func New(financeRepo finance.R, userRepo user.U) finance.U {
	return &financeUsecase{
		financeRepository: financeRepo,
		userUsecase:       userRepo,
	}
}
