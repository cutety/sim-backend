package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type GetApplyMatchingResultService struct {

}

func(*GetApplyMatchingResultService) GetApplyMatchingResult(pagination *common.Pagination, userID string) common.Response {
	result, total, err := models.MUser.GetMatchingResult(pagination, userID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, common.DataList{
		Items: result,
		Total: total,
	})
}