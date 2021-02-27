package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
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


