package financeUsecase

import (
	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/internal/domain/user"
	"github.com/Doittikorn/cypernote/pkg/errorc"
)

func (u *financeUsecase) ExecuteGetByUserID(userId int64, filter *finance.Filter) ([]finance.FinanceResponse, error) {
	// check user id is exist
	var auditUserId user.UserID = userId
	err := u.userUsecase.AuditUserByID(auditUserId)
	var finance = []finance.FinanceResponse{}
	if err != nil {
		return finance, err
	}
	if userId == 0 {
		return finance, errorc.ErrBind
	}

	// get finance by user id
	finance, err = u.financeRepository.GetByUserID(userId, filter.Type)
	if err != nil {
		return finance, err
	}

	return finance, nil
}
