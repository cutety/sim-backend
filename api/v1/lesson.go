package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/service/lesson"
	"sim-backend/utils"
)

// CreateLesson 开始上课
// @Summary 开始上课
// @Tags Lesson
// @Accept json
// @Produce json
// @Param CreateLessonService body lesson.CreateLessonService true "上课信息"
// @Success 200 {object} common.Response
// @Router /mentor/add/lesson [POST]
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
// @Summary 列出未评价的课
// @Tags Lesson
// @Accept json
// @Produce json
// @Param stu_id query string true "学生user_id"
// @Success 200 {object} common.Response
// @Router /student/evaluable/lesson [get]
func ListEvaluableLessons(c *gin.Context) {
	stuID := c.Query("stu_id")
	service := &lesson.ListEvaluableLessons{}
	evaluableLessons, err := service.ListEvaluableLessons(stuID)
	c.JSON(200, utils.ResponseAll(evaluableLessons, err))
}