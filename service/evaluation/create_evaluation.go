package evaluation

import (
	"sim-backend/models"
	"sim-backend/utils"
)

type CreateEvaluationService struct {
	MentorID string `json:"mentor_id" validate:"required" label:"老师ID"`
	LessonID string `json:"lesson_id" validate:"required" label:"课堂ID"`
	CourseID string `json:"course_id" validate:"required" label:"课程ID"`
	StuID string `json:"stu_id" validate:"required" label:"学生ID"`
	Rate string `json:"rate" validate:"required" label:"评分"`
	Content string `json:"content" validate:"required" label:"评价内容"`
}

func (s *CreateEvaluationService) CreateEvaluation() error {
	evaluationID := utils.UUID()
	evaluation := models.Evaluation{
		EvaluationID: evaluationID,
		MentorID: s.MentorID,
		CourseID: s.CourseID,
		LessonID: s.LessonID,
		StuID:    s.StuID,
		Rate:     s.Rate,
		Content:  s.Content,
	}
	return evaluation.Create()
}