package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
	"sim-backend/models/common"
	"time"
)

var MMentor Mentor

type Mentor struct {
	ID                      uint       `gorm:"primary_key"`
	CreatedAt               time.Time  `gorm:"type:timestamp"`
	UpdatedAt               time.Time  `gorm:"type:timestamp"`
	DeletedAt               *time.Time `sql:"index" gorm:"type:timestamp"`
	UserID                  string     `gorm:"column:user_id;type:varchar(20)" json:"user_id" validate:"required" label:"用户ID"`
	Name                    string     `gorm:"column:name;type:varchar(20)" json:"name" validate:"required" label:"姓名"`
	Gender                  int        `gorm:"column:gender;type:int(1);not null;default:1" json:"gender"`
	Phone                   string     `gorm:"column:phone;type:varchar(20);" json:"phone" validate:"required" label:"电话号"`
	Email                   string     `gorm:"column:email;type:varchar(50);" json:"email"`
	Wechat                  string     `grom:"column:wechat;type:varchar(25);" json:"wechat"`
	QQ                      string     `gorm:"column:qq;type:varchar(10);" json:"qq"`
	ResearchDirection       string     `gorm:"column:research_direction;type:varchar(50);" json:"research_direction"`
	Degree                  string     `gorm:"column:degree;type:varchar(25);" json:"degree"`
	UndergraduateUniversity string     `gorm:"column:undergraduate_university;type:varchar(255)" json:"undergraduate_university"`
	UndergraduateMajor      string     `gorm:"column:undergraduate_major;type:varchar(255)" json:"undergraduate_major"`
	GraduateSchool          string     `gorm:"column:graduate_school;type:varchar(255)" json:"graduate_school"`
	GraduateMajor           string     `gorm:"column:graduate_major;type:varchar(255)" json:"graduate_major"`
	PHDSchool               string     `gorm:"phd_school;type:varchar(255)" json:"phd_school"`
	PHDMajor                string     `gorm:"phd_major;type:varchar(255)" json:"phd_major"`
}

func (*Mentor) TableName() string {
	return "mentors"
}

func (*Mentor) Create(mentor *Mentor) error {
	return extension.DB.Create(&mentor).Error
}

func (*Mentor) GetByMentorID(mentorID string) (*Mentor, error) {
	mentor := &Mentor{}
	err := extension.DB.Where("user_id = ?", mentorID).Find(&mentor).Error
	if err == gorm.ErrRecordNotFound {
		return &Mentor{}, err
	}
	if err != nil {
		return nil, err
	}
	return mentor, err
}

type MentorMatchingResult struct {
	UserID            string  `gorm:"column:user_id;type:varchar(20)" json:"user_id" validate:"required" label:"用户ID"`
	StuName           string  `gorm:"column:stu_name;type:varchar(20)" json:"stu_name" validate:"required" label:"姓名"`
	Gender            string  `gorm:"column:gender;type:int(1);not null;default:1" json:"gender"`
	Grade             string  `gorm:"column:grade;type:varchar(4)" json:"grade"`
	Major             string  `gorm:"column:major;type:varchar(50)" json:"major"`
	Phone             string  `gorm:"column:phone;type:varchar(20);" json:"phone" validate:"required" label:"电话号"`
	Email             string  `gorm:"column:email;type:varchar(50);" json:"email"`
	Wechat            string  `grom:"column:wechat;type:varchar(255);" json:"wechat"`
	QQ                string  `gorm:"column:qq;type:varchar(10);" json:"qq"`
	MentorUserID      string  `gorm:"column:mentor_user_id;type:varchar(20)" json:"mentor_user_id" label:"导师用户ID"`
	ApplySchool       string  `gorm:"column:apply_school;type:varchar(255)" json:"apply_school" validate:"required" label:"报考院校"`
	ApplyMajor        string  `gorm:"column:apply_major;type:varchar(255)" json:"apply_major" validate:"required" label:"报考专业"`
	PreliminaryResult float32 `gorm:"column:preliminiary_result;type:decimal(11,2)" json:"preliminiary_result" label:"初试成绩"`
	RetrailResult     float32 `gorm:"column:retrail_result;type:decimal(11,2)" json:"retrail_result" label:"复试成绩"`
	AdmissionSchool   string  `gorm:"column:admission_shcool;type:varchar(255)" json:"admission_shcool" validate:"required" label:"录取院校"`
	AdmissionMajor    string  `gorm:"column:admission_major;type:varchar(255)" json:"admission_major" validate:"required" label:"录取院校"`
	IsAdmitted        bool    `gorm:"column:is_admitted;type:tinyint(1)" json:"is_admitted" label:"录取结果"`
	Note string `gorm:"column:note;" json:"note" label:"留言"`
}

func (*Mentor) GetMatchingResult(pagination *common.Pagination, userID string) ([]MentorMatchingResult, int64, error) {
	var result []MentorMatchingResult
	var total int64
	sql:=`
		SELECT 
			s.stu_name, s.gender, s.grade, s.admission_major as major, s.phone as phone, s.email as email, s.wechat as wechat, s.qq as qq, 
			a.user_id, a.apply_school, a.mentor_user_id, a.apply_school, a.apply_major, a.preliminiary_result, a.retrail_result, a.admission_shcool, a.admission_major, a.is_admitted
		FROM 
			application a 
		left join
			students s
			on s.stu_id = a.user_id 
		left join 
			mentors m 
			on a.apply_school = m.undergraduate_university 
			or a.apply_school = m.graduate_school 
			or a.apply_school = m.phd_school 
		WHERE 
			m.user_id = ?
			AND a.mentor_user_id = ''
			AMD a.status = 1
`

	total = extension.DB.Raw(sql, userID).Scan(&result).RowsAffected
	err := extension.DB.Raw(sql+`
		
			 
    	LIMIT ? OFFSET ? 
		`, userID, pagination.Limit, (pagination.Page-1)*pagination.Limit).
		Scan(&result).Error
	return result, total, err
}

func (*Mentor) ListStudentByMatchingStatus(pagination *common.Pagination, userID string, isMatched int) ([]MentorMatchingResult, int64, error) {
	var apps []MentorMatchingResult
	var total int64
	sql := `
		SELECT
			s.stu_name, s.gender, s.grade, s.admission_major as major, s.phone as phone, s.email as email, s.wechat as wechat, s.qq as qq, 
			a.user_id, a.apply_school, a.mentor_user_id, a.apply_school, a.apply_major, a.preliminiary_result, a.retrail_result, a.admission_shcool, a.admission_major, a.is_admitted, a.note
		FROM
			application a
		left join
			students s
			on s.stu_id = a.user_id
		left join
			mentors m 
			on a.mentor_user_id = m.user_id
		WHERE 
			m.user_id = ?
			AND a.status = ?
`
	total = extension.DB.Raw(sql, userID, isMatched).Scan(&apps).RowsAffected
	err := extension.DB.Raw(sql+`
		LIMIT ? OFFSET ? 
		`, userID, isMatched, pagination.Limit, (pagination.Page-1)*pagination.Limit).
		Scan(&apps).Error
	return apps, total, err
}

func (*Mentor) ListMentors(pagination *common.Pagination) ([]Mentor, int64, error) {
	var apps []Mentor
	var total int64
	err := extension.DB.
		Limit(pagination.Limit).
		Offset((pagination.Page - 1) * pagination.Limit).
		Find(&apps).Error
	extension.DB.Model(&apps).Count(&total)
	return apps, total, err
}

func (*Mentor) UpdateMentorByUserID(userID string, info *Mentor) error {
	return extension.DB.Model(&info).Where("user_id = ?", userID).Updates(&info).Error
}
