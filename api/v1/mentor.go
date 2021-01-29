package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models"
	"sim-backend/models/common"
	"sim-backend/service/mentor"
	"sim-backend/utils"
	"sim-backend/utils/validator"
)

func CreateMemtor(c *gin.Context) {
	var data models.Mentor
	if err := c.ShouldBindJSON(&data); err == nil {
		msg, code := validator.Validate(&data)
		if code != utils.SUCCESS {
			c.JSON(200, common.Response{
				Status: code,
				Msg:    msg,
			})
			c.Abort()
		}
		serivce := mentor.CreateMentorService{}
		response := serivce.CreateMentor(data)
		c.JSON(200, response)
	} else {
		c.JSON(200, utils.Response(utils.ERROR, err))
	}
}
