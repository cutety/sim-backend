package admin

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io"
	"sim-backend/extension"
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/service/user"
	"sim-backend/utils"
	"sim-backend/utils/logger"
)

type BatchAddMentorService struct {
}

type Message struct {
	Status int `json:"status"`
	UserID string `json:"user_id"`
	Msg    string `json:"msg"`
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
	result := make([]Message, len(rows[1:]))
	for index, row := range rows[1:] {
		mentor := &models.Mentor{}
		mentor.UserID = row[14]
		mentor.Name = row[0]
		if row[1] == "女" {
			mentor.Gender = 0
		} else {
			mentor.Gender = 1
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
		createUserService := user.CreateUserService{
			UserID:   mentor.UserID,
			Username: mentor.Name,
			Password: mentor.UserID[len(mentor.UserID)-6:],
			Role:     3,
		}
		err := models.MMentor.Create(mentor)
		if err != nil {
			if extension.IsMySQLDuplicateEntryError(err) {
				result[index].Status = 1
				result[index].UserID = mentor.UserID
				result[index].Msg = "导入失败，重复导入"
			} else {
				result[index].Status = 1
				result[index].UserID = mentor.UserID
				result[index].Msg = fmt.Sprintf("导入失败：%s", err.Error())
			}
		} else {
			_ = createUserService.CreateUser()
			result[index].Status = 0
			result[index].UserID = mentor.UserID
			result[index].Msg = "导入成功"
		}
	}

	return utils.Response(200, result)
}
