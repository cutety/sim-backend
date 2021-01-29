package user

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type LoginService struct {
	UserID string `form:"user_id" json:"user_id"`
	Password string `form:"password" json:"password"`
}

func (service *LoginService) Login() common.Response {
	user, err := models.MUser.GetUserByUserID(service.UserID)
	if err != nil || user == nil{
		return utils.ResponseWithError(utils.ERROR_USER_EXIST, err)
	}
	if !utils.DecodePsw(user.Password, service.Password) {
		return utils.Response(utils.ERROR_PASSWORD_WRONG, nil)
	}
	return utils.Response(utils.SUCCESS, nil)
}