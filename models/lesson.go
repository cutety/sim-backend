package models

import (
	"sim-backend/extension"
	"time"
)

var MLesson Lesson

type Lesson struct {
	ID uint `gorm:"primary_key"`
	CourseID uint `gorm:"course_id;type:uint" json:"course_id"`
	StartAt time.Time `gorm:"start_at;type:timestamp" json:"start_at" validate:"required" label:"上课时间"`
	EndAt time.Time `gorm:"end_at;type:timestamp" json:"end_at" validate:"required" label:"下课时间"`
}

func (*Lesson) TableName() string {
	return "lessons"
}

func (l *Lesson) Create() error{
	result := extension.DB.Where(&l).Attrs(&l).FirstOrCreate(&l)
	if result.RowsAffected != 0 {
		return nil
	} else {
		return result.Error
	}
}

type EvaluableLesson struct {
	LessonID uint `gorm:"column:lesson_id;" json:"lesson_id"`
	StartAt time.Time `gorm:"column:start_at" json:"start_at"`
	EndAt time.Time `gorm:"column:end_at" json:"end_at"`
	Grade string `gorm:"column:grade" json:"grade"`
	Class string `gorm:"column:class" json:"class"`
	Lesson string `gorm:"column:lesson" json:"lesson"`
}

func (*Lesson) ListEvaluableLessons(stuID string) ([]EvaluableLesson, error) {
	var result []EvaluableLesson
	sql := `
		SELECT
			l.id AS lesson_id, l.start_at, l.end_at,
			c.grade, c.class, c.lesson
		FROM
			lessons l
		JOIN 
			courses c
			ON l.course_id = c.id
		LEFT JOIN 
			evaluations e
			ON e.lesson_id = l.id
		WHERE 
			l.id NOT IN
			(
				SELECT
					a.id
				FROM 
				(
					SELECT
						l.id,e.stu_id
					FROM
						lessons l
					LEFT JOIN 
						evaluations e
						ON e.lesson_id = l.id
				) a
				WHERE a.stu_id = ?
			)
		GROUP BY l.id
`
	err := extension.DB.Raw(sql, stuID).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}