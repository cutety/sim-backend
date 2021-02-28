package user

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type UpdateApplicationService struct {
	UserID string `gorm:"column:user_id;type:varchar(20)" json:"user_id" validate:"required" label:"用户ID"`
	MentorUserID string `gorm:"column:mentor_user_id;type:varchar(20)" json:"mentor_user_id" label:"导师用户ID"`
	Status int `gorm:"column:status;" json:"status" label:"匹配情况"`
	Note string `gorm:"column:note;" json:"note" label:"留言"`
}

func(us *UpdateApplicationService) UpdateApplication() common.Response {
	app := &models.Application{
		UserID: us.UserID,
		MentorUserID: us.MentorUserID,
		Status: us.Status,
		Note: us.Note,
	}
	err := models.MApplication.UpdateApplication(app)
	if err != nil {
		return utils.Response(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, nil)
}
