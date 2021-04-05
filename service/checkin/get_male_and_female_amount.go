package checkin

import "sim-backend/models"

type GetMaleAndFemaleAmountService struct {

}

func (s *GetMaleAndFemaleAmountService) GetMaleAndFemaleAmount(grade string) (interface{}, error) {
	result, err := models.MStudent.GetMaleAndFemaleAmount(grade)
	if err != nil {
		return nil, err
	}
	amountMap := map[string]interface{}{}
	for _, item := range result {
		amountMap[item.Gender] = item.Amount
	}
	return amountMap, nil
}