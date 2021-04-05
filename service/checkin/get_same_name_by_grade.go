package checkin

import "sim-backend/models"

type GetSameNameByGradeService struct {

}

func (s *GetSameNameByGradeService) GetSameNameByGrade(grade string) ([]models.StudentsValue, error)  {
	return models.MStudent.GetSameNameByGrade(grade)
}