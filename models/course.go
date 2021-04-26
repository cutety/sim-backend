package models

import (
	"sim-backend/extension"
)

var MCourse Course

type Course struct {
	ID uint `gorm:"primary_key"`
	CourseID string `gorm:"column:course_id;type:varchar(255)" json:"course_id"`
	MentorID string `gorm:"column:mentor_id;type:varchar(20)" json:"mentor_id" validate:"required" label:"老师ID"`
	Lesson string	`gorm:"column:lesson;type:varchar(255)" json:"course" validate:"required" label:"课程名"`
	Grade string 	`gorm:"column:grade;type:varchar(4)" json:"grade" validate:"required" label:"年级"`
	Class string 	`gorm:"column:class;type:varchar(255)" json:"class" validate:"required" label:"班级"`
}

func (*Course) TableName() string {
	return "courses"
}

func (c *Course) Create() error {
	result := extension.DB.Where(&c).Attrs(&c).FirstOrCreate(&c)
	if result.RowsAffected != 0 {
		return nil
	} else {
		return result.Error
	}
}

// ListCourses 通过grade, class获取课程列表
func (*Course) ListCourses(mentorID, grade, class string) ([]Course, error) {
	var result []Course
	err := extension.DB.Where("mentor_id = ? AND grade = ?  AND class = ? ",mentorID, grade, class).
		Find(&result).Error
	return result, err
}