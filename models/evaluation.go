package models

import (
	"sim-backend/extension"
)

var MEvaluation Evaluation

type Evaluation struct {
	ID uint `gorm:"primary_key"`
	MentorID string `gorm:"column:mentor_id;type:varchar(20)" json:"mentor_id" validate:"required" label:"老师ID"`
	CourseID uint `gorm:"column:course_id;type:uint" json:"course_id" validate:"required" label:"课程ID"`
	LessonID uint `gorm:"column:lesson_id;type:uint" json:"lesson_id" validate:"required" label:"课程ID"`
	StuID string `gorm:"column:stu_id;type:varchar(20)" json:"stu_id" validate:"required" label:"学生ID"`
	Rate string `gorm:"column:rate;type:varchar(20)" json:"rate" validate:"required" label:"评分"`
	Content string `gorm:"column:content;type:text" json:"content" validate:"required" label:"评价内容"`
}

func (*Evaluation) TableName() string {
	return "evaluations"
}

func (e *Evaluation) Create() error {
	result := extension.DB.Create(e)
	if result.RowsAffected != 0 {
		return nil
	} else {
		return result.Error
	}
}