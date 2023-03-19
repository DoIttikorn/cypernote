package userUsecase

import (
	"github.com/Doittikorn/cypernote/internal/domain/user"
)

type Usecase struct {
	repository user.R
}

func New(repository user.R) user.U {
	return &Usecase{
		repository: repository,
	}
}

func (u *Usecase) GetByUserID(id user.UserID) (user.M, error) {
	// var t string
	// userID, err := u.Repository.GetByUserID(t)

	// if err != nil {
	// 	return 0, err
	// }

	return user.M{}, nil
}
