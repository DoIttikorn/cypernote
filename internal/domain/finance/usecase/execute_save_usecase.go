package usecase

import (
	"github.com/Doittikorn/cypernote/internal/domain/finance"
	"github.com/Doittikorn/cypernote/internal/domain/user"
)

func (u *financeUsecase) ExecuteSave(request *finance.FinanceRequest) (*finance.FinanceResponse, error) {

	// check user id is exist
	var auditUserId user.UserID = request.UserID
	err := u.userUsecase.AuditUserByID(auditUserId)
	if err != nil {
		return nil, err
	}

	model := finance.M{
		UserID: request.UserID,
		Type:   request.Type,
		Amount: request.Amount,
		Status: request.Status,
	}

	err = u.financeRepository.Save(&model)
	if err != nil {
		return nil, err
	}

	response := finance.FinanceResponse{
		ID:     model.ID,
		Amount: model.Amount,
		Type:   model.Type,
		Note:   model.Note,
		Status: model.Status,
	}

	return &response, nil
}
