package user

import (
	"sim-backend/models"
	"sim-backend/models/common"
)

type GetAdmittedTendencyService struct {

}

func (*GetAdmittedTendencyService) GetAdmittedTendency() ([]common.NameAndValue, error) {
	return models.MApplication.GetAdmittedTendency()
}
