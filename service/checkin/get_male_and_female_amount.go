package checkin

import "sim-backend/models"

type GetMaleAndFemaleAmountService struct {

}

func (s *GetMaleAndFemaleAmountService) GetMaleAndFemaleAmount(grade string) (interface{}, error) {
	female, err := models.MStudent.GetAmountByGender(grade, 0)
	if err != nil {
		return nil, err
	}
	male, err := models.MStudent.GetAmountByGender(grade, 1)
	if err != nil {
		return nil, err
	}
	amountMap := map[string]interface{}{}
	amountMap["male"] = male
	amountMap["female"] = female
	return amountMap, nil
}