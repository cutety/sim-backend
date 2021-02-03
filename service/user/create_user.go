package user

import (
	"sim-backend/extension"
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
	"time"
)

type CreateUserService struct {
	UserID string `gorm:"column:user_id;type:varchar(20)" json:"user_id"`
	Username string `gorm:"column:username;type:varchar(20)" json:"username"`
	Password string `gorm:"column:password;type:varchar(20)" json:"password"`
	Role int `gorm:"column:role;type:int;DEFAULT:2" json:"role"`
}

func (service *CreateUserService) CreateUser() common.Response {

	psw := utils.ScryptPsw(service.Password)
	user := models.User{
		UserID:    service.UserID,
		Username:  service.Username,
		Password:  psw,
		Role:      service.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	err := extension.DB.Create(&user).Error
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, nil)
}