package user

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type GetUserByUserIDService struct {

}

func (*GetUserByUserIDService) GetUserByUserIDService(userID string) common.Response {
	user, err := models.MUser.GetUserByUserID(userID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, user)
}
