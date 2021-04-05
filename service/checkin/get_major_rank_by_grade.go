package checkin

import "sim-backend/models"

type GetMajorRankByGradeService struct {}

func (*GetMajorRankByGradeService) GetMajorRankByGrade(grade string) ([]models.StudentsValue, error) {
	return models.MStudent.GetMajorRankByGrade(grade)
}