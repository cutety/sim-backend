package user

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type UserChangePasswordService struct {

}

func (*UserChangePasswordService) UserChangePassword(id int) common.Response {
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