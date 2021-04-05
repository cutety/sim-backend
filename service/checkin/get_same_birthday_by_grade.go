package checkin

import "sim-backend/models"

type GetSameBirthdayByGradeService struct {}

func (s *GetSameBirthdayByGradeService) GetSameBirthdayByGrade(grade string) ([]models.StudentsValue, error)  {
	return models.MStudent.GetSameBirthdayByGrade(grade)
}
