package usecase

import "github.com/Doittikorn/cypernote/internal/user/repository"

type Usecase struct {
	Repository repository.R
}

type U interface {
	GetByUserID() (float64, error)
}

func New(repository repository.R) U {
	return &Usecase{
		Repository: repository,
	}
}

func (u *Usecase) GetByUserID() (float64, error) {
	// var t string
	// userID, err := u.Repository.GetByUserID(t)

	// if err != nil {
	// 	return 0, err
	// }

	return 0, nil
}
