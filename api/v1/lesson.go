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

// ListEvaluableLessons 列出未评价的课
func ListEvaluableLessons(c *gin.Context) {
	stuID := c.Query("stu_id")
	service := &lesson.ListEvaluableLessons{}
	evaluableLessons, err := service.ListEvaluableLessons(stuID)
	c.JSON(200, utils.ResponseAll(evaluableLessons, err))
}