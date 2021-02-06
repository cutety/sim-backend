package mentor

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type ListMentoredStudentsService struct{}

func (*ListMentoredStudentsService) ListMentoredStudents(userID string) common.Response {
	apps, total, err := models.MMentor.ListMentoredStudents(userID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, common.DataList{
		Items: apps,
		Total: total,
	})
}
