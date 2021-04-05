package checkin

import "sim-backend/models"

type GetAgeDistributionService struct {}

func (s *GetAgeDistributionService) GetAgeDistribution(grade string) (map[string]interface{}, error) {
	result, err := models.MStudent.GetAgeDistribution(grade)
	if err != nil {
		return nil, err
	}
	amountMap := map[string]interface{}{}
	for _, item := range result {
		amountMap[item.Age] = item.Amount
	}
	return amountMap, nil
}
