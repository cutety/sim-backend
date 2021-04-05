package models

import (
	"fmt"
	"sim-backend/extension"
	"time"
)

var MCheckinInfo CheckinInfo

type CheckinInfo struct {
	StuID string `gorm:"stu_id" json:"stu_id"`
	StuName string `gorm:"stu_name" json:"stu_name"`
	Major string `gorm:"major" json:"major"`
	CheckinStatus int `gorm:"checkin_status" json:"checkin_status"`
	CheckinTime time.Time `gorm:"checkin_time" json:"checkin_time"`
}

func (c *CheckinInfo) TableName() string {
	return "checkin_info"
}

func (c *CheckinInfo) UpdateCheckinInfo(checkinInfo *CheckinInfo) error {
	info := map[string]interface{}{
		"stu_id": checkinInfo.StuID,
		"stu_name": checkinInfo.StuName,
		"major": checkinInfo.Major,
		"checkin_status": checkinInfo.CheckinStatus,
		"checkin_time": checkinInfo.CheckinTime,
	}
	return extension.DB.Table(c.TableName()).Where("stu_id = ?", checkinInfo.StuID).Updates(info).Error
}

func (c *CheckinInfo) GetCheckinAmountByGrade(grade string) (int64, error) {
	var total int64
	val := fmt.Sprintf("%s%%", grade)
	err := extension.DB.Table(c.TableName()).Where("stu_id like ? and checkin_status = 1", val).Count(&total).Error
	return total, err
}
