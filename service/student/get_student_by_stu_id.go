package student

import (
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
)

type GetStudentByStuIDService struct{

}

func(*GetStudentByStuIDService) GetStudentByStuID(stuID string) common.Response {
	student, err := models.MStudent.GetByStuID(stuID)
	if err != nil {
		return utils.ResponseWithError(utils.ERROR, err)
	}
	if student == nil {
		return utils.Response(200, &models.Student{})
	}
	return utils.Response(200, student)
}
