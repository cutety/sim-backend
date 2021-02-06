package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type ChooseMentorService struct{

}

func (*ChooseMentorService) ChooseMentor(userID, mentorUserID string) common.Response {
	app, err := models.MApplication.GetByUserID(userID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	if app == nil {
		return utils.Response(utils.ERROR_APPLICATION_EXIST, nil)
	}
	err = models.MApplication.UpdateMentorUserID(userID, mentorUserID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, nil)
}
