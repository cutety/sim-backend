package checkin

import "sim-backend/models"

type GetCheckinAmountService struct {
}

func (s *GetCheckinAmountService) GetCheckinAmount(grade string) (int64, error) {
	return models.MCheckinInfo.GetCheckinAmountByGrade(grade)
}
