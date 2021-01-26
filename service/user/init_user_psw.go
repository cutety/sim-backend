package user

import (
	"sim-backend/models"
	"sim-backend/utils"
)

type InitUserPswService struct {}

func (*InitUserPswService) InitUserPsw() error {
	total, err := models.MUser.Total()
	if err != nil {
		return err
	}
	go func() {
		for i:= 1; i <= *total; i ++ {
			user, _ := models.MUser.GetUserByID(i)
			psw := user.UserID[len(user.UserID) - 6:]
			encodedPsw := utils.ScryptPsw(psw)
			_ = models.MUser.UpdatePasswordById(int(user.ID), encodedPsw)
		}
	}()
	return nil
}
