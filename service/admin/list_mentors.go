package admin

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type ListMentorsService struct {

}

func (*ListMentorsService) ListMentors(pagination *common.Pagination) common.Response {
	mentors, total, err := models.MMentor.ListMentors(pagination)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, common.DataList{
		Items: mentors,
		Total: total,
	})
}