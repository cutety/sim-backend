package lesson

import (
	"sim-backend/models"
	"sim-backend/utils"
	"sim-backend/utils/logger"
)

type CreateLessonService struct {
	CourseID uint `json:"course_id"`
	StartAt string `json:"start_at" validate:"required" label:"上课时间"`
	EndAt string `json:"end_at" validate:"required" label:"下课时间"`
}

func (s *CreateLessonService) CreateLesson() error {
	startAt := utils.ParseWithLocation(s.StartAt)
	endAt := utils.ParseWithLocation(s.EndAt)
	logger.Info("start At:",startAt)
	logger.Info("end At:",endAt)
	lesson := models.Lesson{
		CourseID: s.CourseID,
		StartAt:  startAt,
		EndAt:    endAt,
	}
	return lesson.Create()
}
