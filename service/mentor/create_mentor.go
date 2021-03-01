package mentor

import (
	"sim-backend/extension"
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
	"sim-backend/utils/logger"
)

type CreateMentorService struct {
	UserID string `gorm:"column:user_id;type:varchar(20)" json:"user_id" validate:"required" label:"用户ID"`
	Name string `gorm:"column:name;type:varchar(20)" json:"name" validate:"required" label:"姓名"`
	Gender int `gorm:"column:gender;type:int(1);not null;default:1" json:"gender"`
	Phone string `gorm:"column:phone;type:varchar(20);" json:"phone" validate:"required" label:"电话号"`
	Email string `gorm:"column:email;type:varchar(50);" json:"email"`
	Wechat string `grom:"column:wechat;type:varchar(25);" json:"wechat"`
	QQ string `gorm:"column:qq;type:varchar(10);" json:"qq"`
	ResearchDirection string `gorm:"column:research_direction;type:varchar(50);" json:"research_direction"`
	Degree string `gorm:"column:degree;type:varchar(25);" json:"degree"`
	UndergraduateUniversity string `gorm:"column:undergraduate_university;type:varchar(255)" json:"undergraduate_university"`
	UndergraduateMajor string `gorm:"column:undergraduate_major;type:varchar(255)" json:"undergraduate_major"`
	GraduateSchool string `gorm:"column:graduate_school;type:varchar(255)" json:"graduate_school"`
	GraduateMajor string `gorm:"column:graduate_major;type:varchar(255)" json:"graduate_major"`
	PHDSchool string `gorm:"phd_school;type:varchar(255)" json:"phd_school"`
	PHDMajor string `gorm:"phd_major;type:varchar(255)" json:"phd_major"`
}


func (service *CreateMentorService) CreateMentor() common.Response {
	mentor := &models.Mentor{
		UserID:                  service.UserID,
		Name:                    service.Name,
		Gender:                  service.Gender,
		Phone:                   service.Phone,
		Email:                   service.Email,
		Wechat:                  service.Wechat,
		QQ:                      service.QQ,
		ResearchDirection:       service.ResearchDirection,
		Degree:                  service.Degree,
		UndergraduateUniversity: service.UndergraduateUniversity,
		UndergraduateMajor:      service.UndergraduateMajor,
		GraduateSchool:          service.GraduateSchool,
		GraduateMajor:           service.GraduateMajor,
		PHDSchool:               service.PHDSchool,
		PHDMajor:                service.PHDMajor,
	}
	err := models.MMentor.Create(mentor)
	// 处理Duplicate key error
	if err != nil {
		if extension.IsMySQLDuplicateEntryError(err) {
			logger.Info("Duplicate Key")
			return utils.Response(utils.DUPLICATE_ENTRY_ERROR, err)
		}
		return utils.Response(utils.ERROR, err)
	} else {
		return utils.Response(utils.SUCCESS, nil)
	}
}