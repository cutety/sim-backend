package checkin

import "sim-backend/models"

type GetMaleAndFemaleAmountService struct {

}

func (s *GetMaleAndFemaleAmountService) GetMaleAndFemaleAmount(grade string) (int64, error) {
	return models.MCheckinInfo.GetMaleAndFemaleAmount(grade)
}