package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
	"sim-backend/models/common"
	"time"
)

var MStudent Student

type Student struct {
	ID        uint `gorm:"primary_key"`
	StuID string `gorm:"column:stu_id;type:varchar(20)" json:"stu_id" validate:"required" label:"用户ID"`
	StuName string `gorm:"column:stu_name;type:varchar(20)" json:"stu_name" validate:"required" label:"姓名"`
	Gender string `gorm:"column:gender;type:int(1);not null;default:1" json:"gender"`
	Birthday time.Time `gorm:"column:birthday;type:timestamp" json:"birthday"`
	PolicalStatus string `gorm:"column:polical_status;type:varchar(45)" json:"polical_status"`
	Nation string `gorm:"column:nation;type:varchar(4)" json:"nation"`
	Grade string `gorm:"column:grade;type:varchar(4)" json:"grade"`
	AdmissionMajor string `gorm:"column:admission_major;type:varchar(50)" json:"admission_major"`
	Phone string `gorm:"column:phone;type:varchar(20);" json:"phone" validate:"required" label:"电话号"`
	Email string `gorm:"column:email;type:varchar(50);" json:"email"`
	Wechat string `grom:"column:wechat;type:varchar(25);" json:"wechat"`
	QQ string `gorm:"column:qq;type:varchar(10);" json:"qq"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
	DeletedAt *time.Time `sql:"index" gorm:"type:timestamp"`
}

func (*Student) TableName() string {
	return "students"
}

func (*Student) Update(student *Student) error {
	return extension.DB.Model(&student).Where("stu_id = ?", student.StuID).Updates(&student).Error
}

func (*Student) GetByStuID(stuID string) (*Student, error) {
	stu := &Student{}
	err := extension.DB.Where("stu_id = ?", stuID).Find(&stu).Error
	if err == gorm.ErrRecordNotFound {
		return nil ,nil
	}
	return stu, err
}

type StudentDetail struct {
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
	MentorName string `gorm:"column:mentor_name;type:varchar(20);" json:"mentor_name" label:"导师姓名"`
	ApplySchool       string  `gorm:"column:apply_school;type:varchar(255)" json:"apply_school" validate:"required" label:"报考院校"`
	ApplyMajor        string  `gorm:"column:apply_major;type:varchar(255)" json:"apply_major" validate:"required" label:"报考专业"`
	PreliminaryResult float32 `gorm:"column:preliminiary_result;type:decimal(11,2)" json:"preliminiary_result" label:"初试成绩"`
	RetrailResult     float32 `gorm:"column:retrail_result;type:decimal(11,2)" json:"retrail_result" label:"复试成绩"`
	AdmissionSchool   string  `gorm:"column:admission_shcool;type:varchar(255)" json:"admission_shcool" validate:"required" label:"录取院校"`
	AdmissionMajor    string  `gorm:"column:admission_major;type:varchar(255)" json:"admission_major" validate:"required" label:"录取院校"`
	IsAdmitted        bool    `gorm:"column:is_admitted;type:tinyint(1)" json:"is_admitted" label:"录取结果"`
}

func (s *Student) GetDetailByStuID(pagination *common.Pagination) ([]StudentDetail, int64, error) {
	var total int64
	var apps []StudentDetail
	sql := `
			SELECT 
				s.stu_id as user_id, s.stu_name as stu_name, s.gender as gender, s.grade as grade, s.admission_major as major,
				s.qq as qq, s.phone as phone, s.email as email, s.wechat as wechat, a.mentor_user_id as mentor_user_id, m.name as mentor_name,
				a.apply_school as apply_school, a.apply_major as apply_major, a.preliminiary_result as preliminiary_result,
				a.retrail_result as retrail_result, a.admission_shcool as admission_shcool, a.admission_major as admission_major,
				a.is_admitted as is_admitted
			FROM
				students s
			LEFT JOIN
				application a
					ON a.user_id = s.stu_id
			LEFT JOIN
				mentors m
					ON m.user_id = a.mentor_user_id
`
	total = extension.DB.Raw(sql).Scan(&apps).RowsAffected
	err := extension.DB.Raw(sql+`LIMIT ? OFFSET ?`, pagination.Limit, (pagination.Page - 1) * pagination.Limit).
		Scan(&apps).Error
	if err != nil {
		return nil, 0, err
	}
	return apps, total, err
}

