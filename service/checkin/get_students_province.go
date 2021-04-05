package checkin

import "sim-backend/models"

type GetStudentsProvinceService struct {}

func (s *GetStudentsProvinceService) GetStudentsProvince(grade string)([]models.StudentsValue, error) {
	return models.MStudent.GetStudentsProvince(grade)
}
