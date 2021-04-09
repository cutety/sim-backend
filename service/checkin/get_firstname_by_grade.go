package checkin

import (
	"sim-backend/models"
)

type GetFirstNameByGradeService struct {}

func (s *GetFirstNameByGradeService) GetFirstNameByGrade(grade string) ([]models.StudentsValue, error) {
	topFive, err := models.MStudent.GetFirstnameByGrade(grade, 5)
	if err != nil {
		return nil, err
	}
	all, err := models.MStudent.GetFirstnameByGrade(grade, 0)
	if err != nil {
		return nil, err
	}
	if len(all) > 5 {
		other := models.StudentsValue{
			Name:  "其他",
			Value: int64(len(all) - 5),
		}
		topFive = append(topFive, other)
	}
	return topFive, nil
}
