package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type UpdateInfoService struct{
	StuID string `gorm:"column:stu_id;type:varchar(20)" json:"stu_id" label:"用户ID"`
	Phone string `gorm:"column:phone;type:varchar(20);" json:"phone" label:"电话号"`
	Email string `gorm:"column:email;type:varchar(50);" json:"email"`
	Wechat string `grom:"column:wechat;type:varchar(25);" json:"wechat"`
	QQ string `gorm:"column:qq;type:varchar(10);" json:"qq"`
}

func (us *UpdateInfoService) UpdateInfo() common.Response {
	 info := &models.Student{
	 	StuID: us.StuID,
	 	Phone: us.Phone,
	 	Email: us.Email,
	 	Wechat: us.Wechat,
	 	QQ: us.QQ,
	 }
	err := models.MStudent.Update(info)
	if err != nil {
		return utils.Response(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, nil)
}
