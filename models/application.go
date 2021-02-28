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
	IsAdmitted bool `gorm:"column:is_admitted;" json:"is_admitted" label:"录取结果"`
	Status int `gorm:"column:status;" json:"status" label:"匹配情况"`
	Note string `gorm:"column:note;" json:"note" label:"留言"`
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

func (a *Application) Upsert(info *Application) error {
	user, err := a.GetByUserID(info.UserID)
	if err != nil {
		return err
	}
	if user == nil {
		return a.CreateApplication(info)
	}
	return a.UpdateByUserID(user.UserID, info)
}

func (a *Application) UpdateByUserID(userID string, info *Application) error {
	data := make(map[string]interface{})
	data["mentor_user_id"]=info.MentorUserID
	data["apply_school"]=info.ApplySchool
	data["apply_major"]=info.ApplyMajor
	data["preliminiary_result"]=info.PreliminaryResult
	data["retrail_result"]=info.RetrailResult
	data["admission_shcool"]=info.AdmissionSchool
	data["admission_major"]=info.AdmissionMajor
	data["is_admitted"] = info.IsAdmitted
	data["status"] = info.Status
	data["note"] = info.Note
	return extension.DB.
		Table(a.TableName()).
		Where("user_id = ?", userID).
		Updates(data).
		Error
}

func (*Application) GetByUserID(userID string) (*Application, error) {
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

func (a *Application) UpdateMatchStatus(userID string, status int) error {
	return extension.DB.
		Table(a.TableName()).
		Where("user_id = ?", userID).
		Update("is_matched", status).
		Error
}

func (a *Application) Dissolve(userID string) error {
	info := map[string]interface{}{
		"mentor_user_id":"",
		"status":0,
	}
	return extension.DB.
		Table(a.TableName()).
		Where("user_id = ?", userID).
		Updates(info).
		Error
}

func (a *Application) UpdateApplication(app *Application) error {
	data := map[string]interface{}{
		"user_id":app.UserID,
		"mentor_user_id":app.MentorUserID,
		"status":app.Status,
		"note":app.Note,
	}
	return extension.DB.
		Table(a.TableName()).
		Where("user_id = ?", data["user_id"]).
		Update(data).
		Error
}