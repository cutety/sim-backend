package mentor

import (
	"sim-backend/extension"
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type UpdateMentorService struct {
	UserID string `gorm:"column:user_id;type:varchar(20)" json:"user_id" validate:"required" label:"用户ID"`
	Name string `gorm:"column:name;type:varchar(20)" json:"name" validate:"required" label:"姓名"`
	Gender string `gorm:"column:gender;type:int(1);not null;default:1" json:"gender"`
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

func (service *UpdateMentorService) UpdateMentor() common.Response {
	mentor := models.Mentor{}
	//info := map[string]interface{}{
	//	"name" : service.Name,
	//	"gender": service.Gender,
	//	"phone":                   service.Phone,
	//	"email":                   service.Email,
	//	"wechat":                  service.Wechat,
	//	"qq":                      service.QQ,
	//	"research_direction":       service.ResearchDirection,
	//	"degree":                  service.Degree,
	//	"undergraduate_university": service.UndergraduateUniversity,
	//	"undergraduateMajor":      service.UndergraduateMajor,
	//	"graduate_school":          service.GraduateSchool,
	//	"graduate_major":           service.GraduateMajor,
	//	"phd_school":               service.PHDSchool,
	//	"phd_major":                service.PHDMajor,
	//}
	info := &models.Mentor{
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
	err := extension.DB.Model(&mentor).Where("user_id = ?", service.UserID).Updates(&info).Error
	if err != nil {
		return utils.Response(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, nil)
}
