package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/student"
	"sim-backend/utils"
)

// @Summary 学生更新报考信息
// @Tags Student
// @Accept json
// @Produce json
// @Param student body student.CreateApplicationService true "学生的报考信息"
// @Success 200 {object} common.Response
// @Router /student/application [post]
func CreateApplication(c *gin.Context) {
	var service student.CreateApplicationService
	if err := c.ShouldBindJSON(&service); err == nil {
		response := service.CreateApplication()
		c.JSON(200, response)
	} else {
		c.JSON(200, common.Response{Error: err.Error()})
	}
}

// @Summary 学生获取匹配结果
// @Tags Student
// @Accept json
// @Produce json
// @Param user_id query string false "学生的user_id"
// @Success 200 {object} common.Response
// @Router /user/match/mentor [get]
func GetApplyMatchingResult(c *gin.Context) {
	userID := c.Query("user_id")
	pagination, _ := utils.Pagination(c)
	service := student.GetApplyMatchingResultService{}
	response := service.GetApplyMatchingResult(pagination, userID)
	c.JSON(200, response)
}

// @Summary 学生获取报考信息
// @Tags Student
// @Accept json
// @Produce json
// @Param user_id query string false "学生的user_id"
// @Success 200 {object} common.Response
// @Router /user/apply/info [get]
func GetApplicationInfo(c *gin.Context) {
	userID := c.Query("user_id")
	service := student.GetApplicationInfo{}
	response := service.GetApplicationInfo(userID)
	c.JSON(200, response)
}

// @Summary 学生更新个人信息
// @Tags Student
// @Accept json
// @Produce json
// @Param student body student.UpdateInfoService true "个人信息"
// @Success 200 {object} common.Response
// @Router /student/application [put]
func UpdateInfo(c *gin.Context) {
	var service student.UpdateInfoService
	if err := c.ShouldBindJSON(&service); err == nil {
		response := service.UpdateInfo()
		c.JSON(200, response)
	} else {
		c.JSON(200, common.Response{Error: err.Error()})
	}
}

// @Summary 获取学生信息
// @Tags Student
// @Accept json
// @Produce json
// @Param user_id query string false "学生的user_id"
// @Success 200 {object} common.Response
// @Router /student/info/ [get]
func GetStudent(c *gin.Context) {
	stuID := c.Query("stu_id")
	service := student.GetStudentByStuIDService{}
	response := service.GetStudentByStuID(stuID)
	c.JSON(200, response)
}