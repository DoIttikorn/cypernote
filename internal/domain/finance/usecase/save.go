package usecase

import (
	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/internal/domain/user"
)

func (u *usecase) Save(model *finance.M) error {

	// check user id is exist
	var auditUserId user.UserID = model.UserID
	err := u.userUsecase.AuditUserByID(auditUserId)
	if err != nil {
		return err
	}

	err = u.financeRepository.Save(model)
	if err != nil {
		return err
	}
	return nil
}
