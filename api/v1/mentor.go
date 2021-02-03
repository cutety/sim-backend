package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/mentor"
	"sim-backend/utils"
	"sim-backend/utils/validator"
)

func CreateMemtor(c *gin.Context) {
	var service mentor.CreateMentorService
	if err := c.ShouldBindJSON(&service); err == nil {
		msg, code := validator.Validate(&service)
		if code != utils.SUCCESS {
			c.JSON(200, common.Response{
				Status: code,
				Msg:    msg,
			})
			c.Abort()
			return
		}
		response := service.CreateMentor()
		c.JSON(200, response)
	} else {
		c.JSON(200, utils.Response(utils.ERROR, err))
	}
}

func UpdateMentor(c *gin.Context) {
	var service mentor.UpdateMentorService
	if err := c.ShouldBindJSON(&service); err == nil {
		msg, code := validator.Validate(&service)
		if code != utils.SUCCESS {
			c.JSON(200, common.Response{
				Status: code,
				Msg:    msg,
			})
			c.Abort()
			return
		}
		response := service.UpdateMentor()
		c.JSON(200, response)
	} else {
		c.JSON(200, utils.Response(utils.ERROR, err))
	}
}
