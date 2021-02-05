package user

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type ChooseMentorService struct{

}

func (*ChooseMentorService) ChooseMentor(userID, mentorUserID string) common.Response {
	err := models.MApplication.UpdateMentorUserID(userID, mentorUserID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	return utils.Response(utils.SUCCESS, nil)
}
