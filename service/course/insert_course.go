package course

import (
	"sim-backend/models"
)

type InsertCourseService struct {
	MentorID string `json:"mentor_id" validate:"required" label:"老师ID"`
	Lesson string	`json:"lesson" validate:"required" label:"课程名"`
	Grade string 	`json:"grade" validate:"required" label:"年级"`
	Class string 	`json:"class" validate:"required" label:"班级"`
}

func (s *InsertCourseService) InsertCourse() error {
	course := &models.Course{
		MentorID: s.MentorID,
		Lesson:   s.Lesson,
		Grade:    s.Grade,
		Class:    s.Class,
	}
	return course.Create()
}
