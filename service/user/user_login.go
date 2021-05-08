package user

import (
	"sim-backend/middlewire"
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
	token, response := middlewire.SetToken(user.UserID, user.Role)
	if response.Status != utils.SUCCESS {
		return common.Response{
			Status: response.Status,
			Data:   nil,
			Msg:    response.Msg,
			Error:  response.Error,
		}
	} else {
		info := make(map[string]interface{})
		info["token"] = token
		info["role"] = user.Role
		return common.Response{
			Status: response.Status,
			Data:   info,
			Msg:    response.Msg,
			Error:  response.Error,
		}
	}
}