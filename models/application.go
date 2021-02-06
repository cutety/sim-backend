package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
	"time"
)

var MApplication Application

type Application struct {
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
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
	DeletedAt *time.Time `sql:"index" gorm:"type:timestamp"`
}


func (*Application) TableName() string {
	return "application"
}

func (*Application) CreateApplication(info *Application) error {
	return extension.DB.Where("user_id = ?", info.UserID).Create(&info).Error
}

func (Application) GetByUserID(userID string) (*Application, error) {
	app := &Application{}
	err := extension.DB.Where("user_id = ?", userID).Find(&app).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return app, err
}

func (a *Application) UpdateMentorUserID(userID, mentorUserID string) error {
	return extension.DB.
		Table(a.TableName()).
		Where("user_id = ?", userID).
		Update("mentor_user_id", mentorUserID).
		Error
}