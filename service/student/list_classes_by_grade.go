package student

import "sim-backend/models"

type ListClassesByGradeService struct {}

func (*ListClassesByGradeService) ListClassesByGrade(grade string) ([]string, error) {
	return models.MStudent.ListClassByGrade(grade)
}
