package mentor

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type GetApplyMatchingResultService struct {

}
// @Summary 通过导师ID获取对应的匹配结果
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 object common.Response 成功后返回值
// @Router /mentor/match/{userID} [get]

func(*GetApplyMatchingResultService) GetApplyMatchingResult(pagination *common.Pagination, userID string) common.Response {
	result, total, err := models.MMentor.GetMatchingResult(pagination, userID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, common.DataList{
		Items: result,
		Total: total,
	})
}