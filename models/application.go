package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
	"sim-backend/models/common"
	"time"
)

var MApplication Application

type Application struct {
	ID        uint `gorm:"primary_key"`
	UserID string `gorm:"column:user_id;type:varchar(20)" json:"user_id" validate:"required" label:"用户ID"`
	MentorUserID string `gorm:"column:mentor_user_id;type:varchar(20)" json:"mentor_user_id" label:"导师用户ID"`
	ApplySchool string `gorm:"column:apply_school;type:varchar(255)" json:"apply_school" validate:"required" label:"报考院校"`
	ApplyMajor string `gorm:"column:apply_major;type:varchar(255)" json:"apply_major" validate:"required" label:"报考专业"`
	PreliminaryResult float32 `gorm:"column:preliminary_result;type:decimal(11,2)" json:"preliminary_result" label:"初试成绩"`
	RetrailResult float32 `gorm:"column:retrail_result;type:decimal(11,2)" json:"retrail_result" label:"复试成绩"`
	AdmissionSchool string `gorm:"column:admission_school;type:varchar(255)" json:"admission_school" validate:"required" label:"录取院校"`
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
	data["preliminary_result"]=info.PreliminaryResult
	data["retrail_result"]=info.RetrailResult
	data["admission_school"]=info.AdmissionSchool
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

type MatchedAdmittedStudents struct {
	StuName string `gorm:"stu_name" json:"stu_name"`
	StuID string `gorm:"stu_id" json:"stu_id"`
	Gender string `gorm:"gender" json:"gender"`
	Grade string `gorm:"grade" json:"grade"`
	Major string `gorm:"major" json:"major"`
	Phone string `gorm:"phone" json:"phone"`
	Wechat string `gorm:"wechat" json:"wechat"`
	QQ string `gorm:"qq" json:"qq"`
	Email string `gorm:"email" json:"email"`
	ApplySchool       string  `gorm:"column:apply_school;type:varchar(255)" json:"apply_school" validate:"required" label:"报考院校"`
	ApplyMajor        string  `gorm:"column:apply_major;type:varchar(255)" json:"apply_major" validate:"required" label:"报考专业"`
	PreliminaryResult float32 `gorm:"column:preliminary_result;type:decimal(11,2)" json:"preliminary_result" label:"初试成绩"`
	RetrailResult     float32 `gorm:"column:retrail_result;type:decimal(11,2)" json:"retrail_result" label:"复试成绩"`
	AdmissionSchool string `gorm:"admission_school" json:"admission_school"`
	AdmissionMajor string `gorm:"admission_major" json:"admission_major"`
}

func (*Application) ListMatchedAdmittedStudents(userID string, pagination * common.Pagination) ([]MatchedAdmittedStudents, int64, error) {
	var apps []MatchedAdmittedStudents
	sql := `
	SELECT 
	s.stu_name, s.stu_id, s.gender, s.grade, s.admission_major as major, s.phone, s.wechat, s.email, s.qq, s.email,
	a.admission_school, a.admission_major, a.apply_school, a.apply_major, a.preliminary_result, a.retrail_result
	FROM	
		students s
	LEFT JOIN
		application a
		ON a.user_id = s.stu_id
	JOIN 
		(
			SELECT 
				*
			FROM
			application a
			WHERE a.user_id = ?
		) ap
		ON ap.apply_school = a.admission_school
		AND  a.is_admitted = 1
`
	total := extension.DB.Raw(sql, userID).Scan(&apps).RowsAffected
	paginationSQL := `
	LIMIT ? OFFSET ?
`
	err := extension.DB.Raw(sql+paginationSQL, userID, pagination.Limit, (pagination.Page - 1) * pagination.Limit).Scan(&apps).Error
	if err != nil {
		return nil, 0, err
	}
	return apps, total, err
}

type ApplicationValue struct {
	Admitted  int64 `json:"admitted"`
	Total int64  `json:"total"`
}

func (*Application) GetAdmittedAndNotAdmittedAmount(grade string) ([]ApplicationValue, error) {
	var result []ApplicationValue
	sql := `
		SELECT 
			SUM(case when is_admitted = 1 then 1 ELSE 0 END) AS admitted,
			COUNT(*) AS total
		FROM 
			application a
		LEFT JOIN 
			students s
			ON s.stu_id = a.user_id
		WHERE 
			s.grade = ?
`
	err := extension.DB.Raw(sql, grade).Scan(&result).Error
	return result, err
}