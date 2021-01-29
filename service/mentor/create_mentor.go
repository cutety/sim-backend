package mentor

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type CreateMentorService struct {}


func (*CreateMentorService) CreateMentor(mentor models.Mentor) common.Response {
	err := models.MMentor.Create(mentor)
	if err != nil {
		return utils.Response(utils.ERROR, err)
	} else {
		return utils.Response(utils.SUCCESS, nil)
	}
}