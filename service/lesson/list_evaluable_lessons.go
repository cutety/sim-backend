package lesson

import "sim-backend/models"

type ListEvaluableLessons struct {}

func (*ListEvaluableLessons) ListEvaluableLessons(stuID string) ([]models.EvaluableLesson, error) {
	return models.MLesson.ListEvaluableLessons(stuID)
}
