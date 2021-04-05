package checkin

import "sim-backend/models"

type GetStudentsAmountByGradeService struct {

}

func (s *GetStudentsAmountByGradeService) GetStudentsAmountByGrade(grade string) (int64, error) {
	return models.MStudent.GetStudentsAmountByGrade(grade)
}
