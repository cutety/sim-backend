package admin

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/utils"
	"sim-backend/utils/logger"
)

type BatchAddMentorService struct{

}

func (*BatchAddMentorService) BatchAddMentor(r io.Reader) common.Response {
	f, err := excelize.OpenReader(r)
	if err != nil {
		logger.Error(err)
		return utils.ResponseWithError(utils.ERROR, err)
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		logger.Error(err)
		return utils.ResponseWithError(utils.ERROR, err)
	}
	logger.Info("rows is :", rows)
	for _, row := range rows[1:] {
		mentor := &models.Mentor{}
		mentor.UserID = row[14]
		mentor.Name = row[0]
		if row[1] == "女" {
			mentor.Gender = 0
		} else {
			mentor.Gender= 1
		}
		mentor.Phone = row[2]
		mentor.Email = row[3]
		mentor.Wechat = row[4]
		mentor.QQ = row[5]
		mentor.Degree = row[6]
		mentor.ResearchDirection = row[7]
		mentor.UndergraduateUniversity = row[8]
		mentor.UndergraduateMajor = row[9]
		mentor.GraduateSchool = row[10]
		mentor.GraduateMajor = row[11]
		mentor.PHDSchool = row[12]
		mentor.PHDMajor = row[13]
		err := models.MMentor.Create(mentor)
		if err != nil {
			return utils.ResponseWithError(500, err)
		}
	}
	return utils.Response(200, nil)
}

