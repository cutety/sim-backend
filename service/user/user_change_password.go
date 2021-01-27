package user

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type ChangePasswordService struct {
	ID uint `form:"id" json:"id"`
	UserID string `form:"user_id" json:"user_id"`
	Password string `form:"password" json:"password"`
	NewPassword string `form:"new_password" json:"new_password"`
}

func (service *ChangePasswordService) UserChangePassword(id uint) common.Response {
	user, err := models.MUser.GetUserByID(id)
	if err != nil {
		return common.Response{Error: err.Error()}
	}
	psw := user.UserID[:6]
	encodedPsw := utils.ScryptPsw(psw)
	err = models.MUser.UpdatePasswordById(id, encodedPsw)
	if err != nil {
		return common.Response{Error: err.Error()}
	}
	return common.Response{Msg: "SUC"}
}

func (service *ChangePasswordService) UserChangePasswordByUserID() common.Response {
	user, err := models.MUser.GetUserByUserID(service.UserID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	if !utils.DecodePsw(user.Password, service.Password) {
		return utils.Response(utils.ERROR_PASSWORD_MATCH, nil)
	}
	encodedPsw := utils.ScryptPsw(service.NewPassword)
	err = models.MUser.UpdatePasswordById(user.ID, encodedPsw)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, nil)
}