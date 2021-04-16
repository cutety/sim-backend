package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/service/lesson"
	"sim-backend/utils"
)

// CreateLesson 开始上课
func CreateLesson(c *gin.Context) {
	service := &lesson.CreateLessonService{}
	if err := c.ShouldBindJSON(service); err == nil {
		err := service.CreateLesson()
		c.JSON(200, utils.ResponseAll(nil, err))
	} else {
		c.JSON(200, utils.ResponseAll(nil, err))
	}
}
