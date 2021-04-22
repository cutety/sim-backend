package evaluation

import "sim-backend/models"

type CreateEvaluationService struct {
	MentorID string `json:"mentor_id" validate:"required" label:"老师ID"`
	LessonID uint `json:"lesson_id" validate:"required" label:"课程ID"`
	StuID string `json:"stu_id" validate:"required" label:"学生ID"`
	Rate string `json:"rate" validate:"required" label:"评分"`
	Content string `json:"content" validate:"required" label:"评价内容"`
}

func (s *CreateEvaluationService) CreateEvaluation() error {
	evaluation := models.Evaluation{
		MentorID: s.MentorID,
		LessonID: s.LessonID,
		StuID:    s.StuID,
		Rate:     s.Rate,
		Content:  s.Content,
	}
	return evaluation.Create()
}