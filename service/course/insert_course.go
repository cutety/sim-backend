package course

import (
	"sim-backend/models"
	"sim-backend/utils"
)

type InsertCourseService struct {
	MentorID string `json:"mentor_id" validate:"required" label:"老师ID"`
	Lesson string	`json:"lesson" validate:"required" label:"课程名"`
	Grade string 	`json:"grade" validate:"required" label:"年级"`
	Class string 	`json:"class" validate:"required" label:"班级"`
	ClassID string `json:"class_id" validate:"required" label:"班级ID"`
}

func (s *InsertCourseService) InsertCourse() error {
	courseID := utils.UUID()
	course := &models.Course{
		CourseID: courseID,
		MentorID: s.MentorID,
		Lesson:   s.Lesson,
		Grade:    s.Grade,
		Class:    s.Class,
		ClassID: s.ClassID,
	}
	return course.Create()
}
