package checkin

import "sim-backend/models"

type GetCheckinInfoByGradeService struct {

}

func (*GetCheckinInfoByGradeService) GetCheckinInfoByGrade(grade string) ([]models.CheckinTable, error) {
	return models.MCheckinInfo.GetCheckinInfoByGrade(grade)
}