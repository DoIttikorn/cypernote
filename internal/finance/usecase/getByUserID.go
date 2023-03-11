package usecase

import (
	"github.com/Doittikorn/cypernote/internal/finance"
	"github.com/Doittikorn/cypernote/internal/user"
	"github.com/Doittikorn/cypernote/pkg/errorc"
)

func (u *usecase) GetByUserID(userId float64, filter *finance.Filter) ([]finance.M, error) {
	// check user id is exist
	var auditUserId user.UserID = userId
	err := u.UserRepository.AuditUserByID(auditUserId)
	var finance = []finance.M{}
	if err != nil {
		return finance, err
	}
	if userId == 0 {
		return finance, errorc.ErrBind
	}

	// get finance by user id
	finance, err = u.FinanceRepository.GetByUserID(userId, filter.Type)
	if err != nil {
		return finance, err
	}

	return finance, nil
}
