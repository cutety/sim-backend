package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type GetApplicationInfo struct {

}

type ApplicationInfo struct {
	ID        uint `gorm:"primary_key"`
	UserID string `gorm:"column:user_id;type:varchar(20)" json:"user_id" validate:"required" label:"用户ID"`
	MentorUserID string `gorm:"column:mentor_user_id;type:varchar(20)" json:"mentor_user_id" label:"导师用户ID"`
	ApplySchool string `gorm:"column:apply_school;type:varchar(255)" json:"apply_school" validate:"required" label:"报考院校"`
	ApplyMajor string `gorm:"column:apply_major;type:varchar(255)" json:"apply_major" validate:"required" label:"报考专业"`
	PreliminaryResult float32 `gorm:"column:preliminiary_result;type:decimal(11,2)" json:"preliminiary_result" label:"初试成绩"`
	RetrailResult float32 `gorm:"column:retrail_result;type:decimal(11,2)" json:"retrail_result" label:"复试成绩"`
	AdmissionSchool string `gorm:"column:admission_shcool;type:varchar(255)" json:"admission_shcool" validate:"required" label:"录取院校"`
	AdmissionMajor string `gorm:"column:admission_major;type:varchar(255)" json:"admission_major" validate:"required" label:"录取院校"`
	IsAdmitted bool `gorm:"column:is_admitted;type:tinyint(1)" json:"is_admitted" label:"录取结果"`
	MentorName string `gorm:"column:name;type:varchar(20)" json:"mentor_name" validate:"required" label:"导师姓名"`
}

func (*GetApplicationInfo) GetApplicationInfo(userID string) common.Response {
	application, err := models.MApplication.GetByUserID(userID)
	if err != nil{
		return utils.ResponseWithError(utils.ERROR, err)
	}
	if application == nil {
		return utils.Response(200, &ApplicationInfo{})
	}
	mentor := &models.Mentor{}
	if application.MentorUserID != "" {
		mentor, err = models.MMentor.GetByMentorID(application.MentorUserID)
		if err != nil || mentor == nil {
			return utils.ResponseWithError(utils.ERROR, err)
		}
	}
	return utils.Response(200, ApplicationInfo{
		ID: application.ID,
		UserID: application.UserID,
		MentorUserID: application.MentorUserID,
		ApplySchool: application.ApplySchool,
		ApplyMajor: application.ApplyMajor,
		PreliminaryResult: application.PreliminaryResult,
		RetrailResult:application.RetrailResult,
		AdmissionSchool:application.AdmissionSchool,
		AdmissionMajor:application.AdmissionMajor,
		IsAdmitted :application.IsAdmitted,
		MentorName:mentor.Name,
	})
}