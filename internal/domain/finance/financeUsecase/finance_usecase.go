package financeUsecase

import (
	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/internal/domain/user"
)

// don't another package can access this struct
// and property can't access from another package
type financeUsecase struct {
	financeRepository finance.R
	userUsecase       user.U
}

// initialize usecase for finance when call New Usecase can use
// finance repository and user repository
func New(financeRepo finance.R, userUsecase user.U) finance.U {
	return &financeUsecase{
		financeRepository: financeRepo,
		userUsecase:       userUsecase,
	}
}
