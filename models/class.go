package models

import (
	"sim-backend/extension"
	"time"
)

var MClass Class

type Class struct {
	ID uint `gorm:"primary_key"`
	ClassID string `gorm:"column:class_id;type:varchar(255)" json:"class_id"`
	Grade string `gorm:"column:grade;type:varchar(50)" json:"grade"`
	Name string `gorm:"column:name;type:varchar(50)" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (*Class) TableName() string {
	return "class"
}

func (c *Class) Create() error {
	result := extension.DB.Where(&c).Attrs(&c).FirstOrCreate(&c)
	if result.RowsAffected != 0 {
		return nil
	} else {
		return result.Error
	}
}

// ListClassesByGrade 通过年级获取班级列表
func (*Class) ListClassesByGrade(grade string) ([]Class, error){
	var result []Class
	err := extension.DB.Where("grade = ?", grade).Find(&result).Error
	return result, err
}