package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/course"
	"sim-backend/utils"
	"sim-backend/utils/validator"
)

func InsertCourse(c *gin.Context) {
	service := &course.InsertCourseService{}
	if err := c.ShouldBindJSON(service); err == nil {
		msg, code := validator.Validate(service)
		if code != utils.SUCCESS {
			c.JSON(200, common.Response{
				Status: code,
				Msg:    msg,
			})
			c.Abort()
			return
		}
		err := service.InsertCourse()
		c.JSON(200, utils.ResponseAll(nil, err))
	} else {
		c.JSON(utils.ERROR, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
	}
}
