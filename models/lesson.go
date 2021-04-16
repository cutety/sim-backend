package models

import (
	"sim-backend/extension"
	"time"
)

type Lesson struct {
	ID uint `gorm:"primary_key"`
	CourseID uint `gorm:"course_id;type:uint" json:"course_id"`
	StartAt time.Time `gorm:"start_at;type:timestamp" json:"start_at" validate:"required" label:"上课时间"`
	EndAt time.Time `gorm:"end_at;type:timestamp" json:"end_at" validate:"required" label:"下课时间"`
}

func (*Lesson) TableName() string {
	return "lesson"
}

func (l *Lesson) Create() error{
	result := extension.DB.Where(&l).Attrs(&l).FirstOrCreate(&l)
	if result.RowsAffected != 0 {
		return nil
	} else {
		return result.Error
	}
}
