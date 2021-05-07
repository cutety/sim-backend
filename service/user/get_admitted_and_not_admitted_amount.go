package user

import "sim-backend/models"

type GetAdmittedAndNotAdmittedAmountService struct{}

func (*GetAdmittedAndNotAdmittedAmountService) GetAdmittedAndNotAdmittedAmount(grade string) ([]models.ApplicationValue, error) {
	return models.MApplication.GetAdmittedAndNotAdmittedAmount(grade)
}
