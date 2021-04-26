package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/course"
	"sim-backend/utils"
	"sim-backend/utils/validator"
)


// InsertCourse 插入课程
// @Summary 开始上课
// @Tags Course
// @Accept json
// @Produce json
// @Param InsertCourseService body course.InsertCourseService true "评价信息"
// @Success 200 {object} common.Response
// @Router /mentor/add/course [POST]
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

func ListCourses(c *gin.Context) {
	mentorID := c.Query("mentor_id")
	grade := c.Query("grade")
	class := c.Query("class")
	service := &course.ListCoursesService{}
	result, err := service.ListService(mentorID, grade, class)
	c.JSON(200, utils.ResponseAll(result, err))
}