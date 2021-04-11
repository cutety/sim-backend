package checkin

import "sim-backend/models"

type GetStudentsInfoTableService struct {}

func (s *GetStudentsInfoTableService) GetStudentsInfoTable(grade string) ([]models.StudentsInfoTable, error) {
	return models.MStudent.GetStudentsInfoTable(grade)
}