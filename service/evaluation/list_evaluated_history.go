package evaluation

import (
	"sim-backend/models"
)

type ListEvaluatedHistoryService struct {}

func (*ListEvaluatedHistoryService) ListEvaluatedHistory(stuID string) ([]models.EvaluatedHistory, int64, error) {
	return models.MEvaluation.ListEvaluatedHistory(stuID)
}
