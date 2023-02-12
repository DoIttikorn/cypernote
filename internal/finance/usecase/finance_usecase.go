package usecase

import (
	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/internal/finance/repository"
)

type Usecase struct {
	Repository repository.Repository
}

type FinanceUsecase interface {
	GetByUserID()
	Save(model *finance.M) error
	Update()
}

func New(repository repository.Repository) FinanceUsecase {
	return &Usecase{
		Repository: repository,
	}
}

func (u *Usecase) GetByUserID() {
}

func (u *Usecase) Save(model *finance.M) error {
	err := u.Repository.Save(model)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) Update() {
}
