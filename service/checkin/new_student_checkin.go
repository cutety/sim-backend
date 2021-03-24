package checkin

import (
	"sim-backend/models"
	"time"
)

type NewStudentCheckinService struct {
	StuID string `json:"stu_id" validate:"required" label:"学号"`
	StuName string `json:"stu_name" validate:"required" label:"姓名"`
	Major string `json:"major" validate:"required" label:"专业"`
}

func (c *NewStudentCheckinService) CheckIn() error {
	info := &models.CheckinInfo{
		StuID:         c.StuID,
		StuName:       c.StuName,
		Major:         c.Major,
		CheckinStatus: 1,
		CheckinTime:   time.Now(),
	}
	return models.MCheckinInfo.UpdateCheckinInfo(info)
}

