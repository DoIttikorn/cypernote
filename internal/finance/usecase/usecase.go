package usecase

import (
	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/internal/finance/repository"
	userRepository "github.com/Doittikorn/cypernote/internal/user/repository"
)

type Usecase struct {
	FinanceRepository repository.R
	UserRepository    userRepository.R
}

type U interface {
	GetByUserID(userID float64, filter *finance.Filter) ([]finance.M, error)
	Save(model *finance.M) error
	Update()
}

func New(financeRepo repository.R, userRepo userRepository.R) U {
	return &Usecase{
		FinanceRepository: financeRepo,
		UserRepository:    userRepo,
	}
}

func (u *Usecase) Update() {
}
