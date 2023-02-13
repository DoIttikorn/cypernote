package usecase

import (
	"time"

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
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	model.CreatedAt = currentTime

	err := u.Repository.Save(model)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) Update() {
}
