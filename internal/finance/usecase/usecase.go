package usecase

import (
	"github.com/Doittikorn/cypernote/internal/finance"
	userRepository "github.com/Doittikorn/cypernote/internal/user/repository"
)

type Usecase struct {
	FinanceRepository finance.R
	UserRepository    userRepository.R
}

func New(financeRepo finance.R, userRepo userRepository.R) finance.U {
	return &Usecase{
		FinanceRepository: financeRepo,
		UserRepository:    userRepo,
	}
}

func (u *Usecase) Update() {
}
