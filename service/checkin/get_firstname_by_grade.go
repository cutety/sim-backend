package checkin

import "sim-backend/models"

type GetFirstNameByGradeService struct {}

func (s *GetFirstNameByGradeService) GetFirstNameByGrade(grade string) ([]models.StudentsValue, error) {
	return models.MStudent.GetFirstnameByGrade(grade)
}
