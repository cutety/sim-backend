package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
	"time"
)

type CreateApplicationService struct {
	UserID            string  `form:"user_id" json:"user_id"`
	MentorUserID      string  `gorm:"column:mentor_user_id;type:varchar(20)" json:"mentor_user_id" label:"导师用户ID"`
	ApplySchool       string  `form:"apply_school" json:"apply_school"`
	ApplyMajor        string  `form:"apply_major" json:"apply_major" `
	PreliminaryResult float32 `form:"preliminiary_result" json:"preliminiary_result"`
	RetrailResult     float32 `form:"retrail_result" json:"retrail_result"`
	AdmissionSchool   string  `form:"admission_shcool" json:"admission_shcool"`
	AdmissionMajor    string  `form:"admission_major" json:"admission_major"`
	IsAdmitted        bool    `form:"is_admitted" json:"is_admitted"`
	Status            int     `gorm:"column:status;" json:"status" label:"匹配情况"`
	Note              string  `gorm:"column:note;" json:"note" label:"留言"`
}

func (service *CreateApplicationService) CreateApplication() common.Response {
	application := &models.Application{
		UserID:            service.UserID,
		MentorUserID:      service.MentorUserID,
		ApplySchool:       service.ApplySchool,
		ApplyMajor:        service.ApplyMajor,
		PreliminaryResult: service.PreliminaryResult,
		RetrailResult:     service.RetrailResult,
		AdmissionSchool:   service.AdmissionSchool,
		AdmissionMajor:    service.AdmissionMajor,
		IsAdmitted:        service.IsAdmitted,
		Status:            service.Status,
		Note:              service.Note,
		CreatedAt:         time.Time{},
		UpdatedAt:         time.Time{},
		DeletedAt:         nil,
	}
	err := models.MApplication.Upsert(application)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, nil)
}
