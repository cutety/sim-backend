package mentor

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type ListMentoredStudentsService struct{}

func (*ListMentoredStudentsService) ListMentoredStudents(pagination *common.Pagination, userID string, status int) common.Response {
	apps, total, err := models.MMentor.ListStudentByMatchingStatus(pagination, userID, status)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, common.DataList{
		Items: apps,
		Total: total,
	})
}
