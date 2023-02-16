package usecase

import (
	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/internal/finance/repository"
	UserRepository "github.com/Doittikorn/cypernote/internal/user/repository"
	"github.com/Doittikorn/cypernote/pkg/errorc"
)

type Usecase struct {
	FinanceRepository repository.R
	UserRepository    UserRepository.R
}

type U interface {
	GetByUserID(userID float64) ([]finance.M, error)
	Save(model *finance.M) error
	Update()
}

func New(financeRepo repository.R, userRepo UserRepository.R) U {
	return &Usecase{
		FinanceRepository: financeRepo,
		UserRepository:    userRepo,
	}
}

func (u *Usecase) GetByUserID(userId float64) ([]finance.M, error) {
	userId, err := u.UserRepository.GetByUserID(userId)
	var finance = []finance.M{}
	if err != nil {
		return finance, err
	}
	if userId == 0 {
		return finance, errorc.New("user not found")
	}

	finance, err = u.FinanceRepository.GetByUserID(userId)
	if err != nil {
		return finance, err
	}

	return finance, nil

}

func (u *Usecase) Save(model *finance.M) error {

	err := u.FinanceRepository.Save(model)
	if err != nil {
		return err
	}
	return nil
}

func (u *Usecase) Update() {
}
