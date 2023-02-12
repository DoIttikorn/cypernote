package usecase

import (
	"github.com/Doittikorn/cypernote/internal/finance/repository"
)

type Usecase struct {
	Repository repository.Repository
}

type FinanceUsecase interface {
	GetByUserID()
	Save()
	Update()
}

func New(repository repository.Repository) FinanceUsecase {
	return &Usecase{
		Repository: repository,
	}
}

func (u *Usecase) GetByUserID() {
}

func (u *Usecase) Save() {
}

func (u *Usecase) Update() {
}
