package evaluation

import "sim-backend/models"

type ListEvaluationService struct {

}

func (*ListEvaluationService) ListEvaluation(mentorID, courseID string) ([]models.EvaluationDetail, int64, error) {
	return models.MEvaluation.ListEvaluation(mentorID, courseID)
}
