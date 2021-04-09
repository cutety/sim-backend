package checkin

import "sim-backend/models"

type GetAgeDistributionService struct {}

func (s *GetAgeDistributionService) GetAgeDistribution(grade string) ([]models.StudentsValue, error) {
	result, err := models.MStudent.GetAgeDistribution(grade)
	if err != nil {
		return nil, err
	}
	return result, nil
}
