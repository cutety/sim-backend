package user

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type DissolveService struct {

}

func (*DissolveService) Dissolve(userID string) common.Response {
	err := models.MApplication.Dissolve(userID)
	if err != nil {
		return utils.ResponseWithError(500, err)
	}
	return utils.Response(200, nil)
}