package usecase

import "github.com/Doittikorn/cypernote/internal/finance"

func (u *Usecase) Save(model *finance.M) error {

	err := u.FinanceRepository.Save(model)
	if err != nil {
		return err
	}
	return nil
}
