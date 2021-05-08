package models

import (
	"sim-backend/extension"
	"time"
)

var MLesson Lesson

type Lesson struct {
	ID uint `gorm:"primary_key"`
	LessonID string `gorm:"lesson_id;type:varchar(255)" json:"lesson_id"`
	CourseID string `gorm:"course_id;type:varchar(255)" json:"course_id"`
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
	LessonID string `gorm:"column:lesson_id;" json:"lesson_id"`
	CourseID string `gorm:"column:course_id" json:"course_id"`
	Grade string `gorm:"column:grade" json:"grade"`
	Class string `gorm:"column:class" json:"class"`
	Lesson string `gorm:"column:lesson" json:"lesson"`
	MentorID string `gorm:"column:mentor_id" json:"mentor_id"`
	MentorName string `gorm:"column:mentor_name" json:"mentor_name"`
	StartAt time.Time `gorm:"column:start_at" json:"start_at"`
	EndAt time.Time `gorm:"column:end_at" json:"end_at"`
}

func (*Lesson) ListEvaluableLessons(stuID string) ([]EvaluableLesson, error) {
	var result []EvaluableLesson
	sql := `
SELECT 
	l.lesson_id, ANY_VALUE(l.start_at) AS start_at, ANY_VALUE(l.end_at) AS end_at,
	ANY_VALUE(c.course_id) as course_id, ANY_VALUE(c.class_id) as class_id, ANY_VALUE(c.grade) AS grade, ANY_VALUE(c.class) AS class, ANY_VALUE(c.lesson) AS lesson,
	ANY_VALUE(c.mentor_id) AS mentor_id, ANY_VALUE(m.name) AS mentor_name
FROM 
	lessons l
JOIN
	courses c
	ON l.course_id = c.course_id
LEFT JOIN 
	evaluations e
	ON e.lesson_id = l.lesson_id
LEFT JOIN 
	mentors m
	ON m.user_id = c.mentor_id
WHERE
	l.lesson_id NOT IN
(
	SELECT
		a.lesson_id
	FROM
	(
		SELECT
			l.lesson_id, e.stu_id
		FROM 
			lessons l 
		LEFT JOIN 
			evaluations e
			ON e.lesson_id = l.lesson_id
		WHERE e.stu_id = ?
	) a 
)
AND c.class_id = (
						SELECT 
							class_id
						FROM 
							class
						WHERE 
							grade = (SELECT grade FROM students WHERE stu_id = ?)
							AND name = (SELECT class FROM students WHERE stu_id = ?)
						)
GROUP BY lesson_id
`
	err := extension.DB.Raw(sql, stuID, stuID, stuID).Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}