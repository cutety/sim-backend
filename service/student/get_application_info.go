package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type GetApplicationInfo struct {

}

func (*GetApplicationInfo) GetApplicationInfo(userID string) common.Response {
	application, err := models.MApplication.GetByUserID(userID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(200, application)
}