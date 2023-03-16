package usecase

import "github.com/Doittikorn/cypernote/internal/domain/user"

func (u *Usecase) AuditUserByID(id user.UserID) error {

	err := u.repository.AuditUserByID(id)

	if err != nil {
		return err
	}

	return nil
}
