package evaluation

import "sim-backend/models"

type GetEvaluationDetailService struct {}

func (*GetEvaluationDetailService) GetEvaluationDetail(evaluationID string) (*models.Evaluation, error) {
	return models.MEvaluation.GetEvaluationDetail(evaluationID)
}
