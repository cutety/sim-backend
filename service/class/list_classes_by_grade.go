package class

import "sim-backend/models"

type ListClassesByGradeService struct {

}

func (*ListClassesByGradeService) ListClassesByGrade(grade string) ([]models.Class, error) {
	return models.MClass.ListClassesByGrade(grade)
}
