package evaluation

type CreateEvaluationService struct {
	MentorID string `json:"mentor_id" validate:"required" label:"老师ID"`
	CourseID uint `json:"course_id" validate:"required" label:"课程ID"`
	StuID string `json:"stu_id" validate:"required" label:"学生ID"`
	Rate string `json:"rate" validate:"required" label:"评分"`
	Content string `json:"content" validate:"required" label:"评价内容"`
}
