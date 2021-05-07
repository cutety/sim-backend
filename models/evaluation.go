package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
	"time"
)

var MEvaluation Evaluation

type Evaluation struct {
	ID uint `gorm:"primary_key"`
	EvaluationID string `gorm:"column:evaluation_id;type:varchar(255)" json:"evaluation_id"`
	MentorID string `gorm:"column:mentor_id;type:varchar(20)" json:"mentor_id" validate:"required" label:"老师ID"`
	CourseID string `gorm:"column:course_id;type:varchar(255)" json:"course_id" validate:"required" label:"课程ID"`
	LessonID string `gorm:"column:lesson_id;type:varchar(255)" json:"lesson_id" validate:"required" label:"课堂ID"`
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

type EvaluationDetail struct {
	EvaluationID string `gorm:"evaluation_id" json:"evaluation_id"`
	Lesson string `gorm:"lesson" json:"lesson"`
	LessonID string `gorm:"lesson_id" json:"lesson_id"`
	StuName string `gorm:"stu_name" json:"stu_name"`
	StuID string `gorm:"stu_id" json:"stu_id"`
	Content string `gorm:"content" json:"content"`
	Rate string `gorm:"rate" json:"rate"`
	StartAt time.Time `gorm:"start_at" json:"start_at"`
	EndAt time.Time `gorm:"end_at" json:"end_at"`
}

func (*Evaluation) ListEvaluation(mentorID, courseID string) ([]EvaluationDetail, int64, error) {
	var result []EvaluationDetail
	var total int64
	sql := `
		SELECT 
			e.evaluation_id, c.lesson, e.lesson_id, e.stu_id, s.stu_name ,e.content, e.rate, l.start_at, l.end_at
		FROM 
			evaluations e
		LEFT JOIN 
			lessons l
			ON l.lesson_id = e.lesson_id
		LEFT JOIN 
			courses c
			ON c.course_id = l.course_id
		LEFT JOIN
			students s
			ON s.stu_id = e.stu_id
	`
	where := `WHERE e.mentor_id = ?`
	var whereVals []interface{}
	whereVals = append(whereVals, mentorID)
	if courseID != "" {
		where += ` AND c.course_id = ?`
		whereVals = append(whereVals, courseID)
	}
	//err := extension.DB.Where(where, whereVals...).
	//	Find(&result).Count(&total).Error
	err := extension.DB.Raw(sql+where, whereVals...).Scan(&result).Error
	total = extension.DB.Raw(sql+where, whereVals...).Scan(&result).RowsAffected
	return result, total, err
}

// GetEvaluationDetail 根据evaluationID获取课程评价详情
func (* Evaluation) GetEvaluationDetail(evaluationID string) (*Evaluation, error) {
	evaluation := &Evaluation{}
	err := extension.DB.Where("evaluation_id = ?", evaluationID).Find(evaluation).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return evaluation, nil
	}
	if err != nil {
		return evaluation, err
	}
	return evaluation, nil
}

type EvaluatedHistory struct {
	EvaluationID string `gorm:"evaluation_id" json:"evaluation_id"`
	Lesson string `gorm:"lesson" json:"lesson"`
	LessonID string `gorm:"lesson_id" json:"lesson_id"`
	Name string `gorm:"name" json:"name"`
	MentorID string `gorm:"mentor_id" json:"mentor_id"`
	Content string `gorm:"content" json:"content"`
	Rate string `gorm:"rate" json:"rate"`
	StartAt time.Time `gorm:"start_at" json:"start_at"`
	EndAt time.Time `gorm:"end_at" json:"end_at"`
}

func (*Evaluation) ListEvaluatedHistory(stuID string) ([]EvaluatedHistory, int64, error) {
	var result []EvaluatedHistory
	var total int64
	sql := `
		SELECT 
			c.lesson, e.mentor_id,m.name, e.content, e.rate, l.start_at, l.end_at
		FROM 
			evaluations e
		LEFT JOIN 
			lessons l
			ON l.lesson_id = e.lesson_id
		LEFT JOIN 
			courses c
			ON c.course_id = l.course_id
		LEFT JOIN
			mentors m
			ON m.user_id = e.mentor_id
		WHERE 
			e.stu_id = ?
	`
	err := extension.DB.Raw(sql, stuID).Scan(&result).Error
	total = extension.DB.Raw(sql, stuID).Scan(&result).RowsAffected
	return result, total, err

}