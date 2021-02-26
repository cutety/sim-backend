package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
	"time"
)

type CreateApplicationService struct {
	UserID string `form:"user_id" json:"user_id"`
	ApplySchool string `form:"apply_school" json:"apply_school"`
	ApplyMajor string `form:"apply_major" json:"apply_major" `
	PreliminaryResult float32 `form:"preliminiary_result" json:"preliminiary_result"`
	RetrailResult float32 `form:"retrail_result" json:"retrail_result"`
	AdmissionSchool string `form:"admission_shcool" json:"admission_shcool"`
	AdmissionMajor string `form:"admission_major" json:"admission_major"`
	IsAdmitted bool `form:"is_admitted" is_admitted:"is_admitted"`
}

func (service *CreateApplicationService) CreateApplication() common.Response {
	application := &models.Application{
		UserID:            service.UserID,
		ApplySchool:       service.ApplySchool,
		ApplyMajor:        service.ApplyMajor,
		PreliminaryResult: service.PreliminaryResult,
		RetrailResult:     service.RetrailResult,
		AdmissionSchool:   service.AdmissionSchool,
		AdmissionMajor:    service.AdmissionMajor,
		IsAdmitted:        service.IsAdmitted,
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