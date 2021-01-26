package user

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
	"sim-backend/utils/logger"
)

type UserLoginService struct {
	UserID string `form:"user_id" json:"user_id"`
	Password string `form:"password" json:"password"`
}

func (service *UserLoginService) Login() common.Response {
	logger.Info(service.UserID[len(service.UserID) - 6:])
	user, err := models.MUser.GetUserByUserID(service.UserID)
	if err != nil || user == nil{
		return common.Response{Msg: "账号不存在"}
	}

	if !utils.DecodePsw(user.Password, service.Password) {
		return common.Response{Msg: "密码不正确"}
	}
	return common.Response{Msg: "登陆成功"}
}