package mentor

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type GetMentorInfoService struct {

}

func (*GetMentorInfoService) GetMentorInfo(userID string) common.Response  {
	mentor, err := models.MMentor.GetByMentorID(userID)
	if err != nil || mentor == nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(200, mentor)
}