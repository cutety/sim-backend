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
	var result []int64
	sql := `
		SELECT 
			COUNT(*) AS amount
		FROM 
			checkin_info c
		LEFT JOIN 
			students s
			ON s.stu_id = c.stu_id
		WHERE 
			s.grade = ?
			AND c.checkin_status = 1
`
	err := extension.DB.Raw(sql, grade).Pluck("amount", &result).Error
	fmt.Println(result)
	return result[0], err
}

type CheckinTable struct {
	StuName string `gorm:"stu_name" json:"stu_name"`
	Major string `gorm:"major" json:"major"`
	CheckinTime time.Time `gorm:"checkin_time" json:"checkin_time"`
}

// GetCheckinInfoByGrade 通过年级获取报到信息
func (c *CheckinInfo) GetCheckinInfoByGrade(grade string) ([]CheckinTable, error) {
	var app []CheckinTable
	err := extension.DB.Raw(`
	SELECT 
		c.stu_name,
		c.major,
		c.checkin_time
	FROM 
		checkin_info c
	LEFT JOIN 
		students as s
	ON 
		c.stu_id = s.stu_id
	WHERE 
		grade = ?
		AND c.checkin_status = 1
	`, grade).Scan(&app).Error
	return app, err
}
