package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
)

type ListMatchedAdmittedStudentsService struct {

}

func (*ListMatchedAdmittedStudentsService) ListMatchedAdmittedStudents(userID string, pagination *common.Pagination) ([]models.MatchedAdmittedStudents, int64, error) {
	students, total, err := models.MApplication.ListMatchedAdmittedStudents(userID, pagination)
	if err != nil {
		return nil, 0, err
	}
	return students, total, err
}
