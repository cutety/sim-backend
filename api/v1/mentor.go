package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/mentor"
	"sim-backend/utils"
	"sim-backend/utils/validator"
)

// @Summary 添加导师
// @Tags Mentor
// @Accept json
// @Produce json
// @Param mentor body mentor.CreateMentorService true "导师的个人信息"
// @Success 200 {object} common.Response
// @Router /mentor [post]
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

// @Summary 更新导师信息
// @Tags Mentor
// @Accept json
// @Produce json
// @Param mentor body mentor.UpdateMentorService true "导师的个人信息"
// @Success 200 {object} common.Response
// @Router /admin/mentor [post]
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

// @Summary 导师根据userID获取匹配结果
// @Tags Mentor
// @Accept json
// @Produce json
// @Param user_id query string false "导师的user_id"
// @Success 200 {object} common.Response
// @Router /mentor/match [get]
func GetMentorMatchingResult(c *gin.Context) {
	userID := c.Query("user_id")
	pagination, _ := utils.Pagination(c)
	service := mentor.GetApplyMatchingResultService{}
	response := service.GetApplyMatchingResult(pagination, userID)
	c.JSON(200, response)
}

// @Summary 导师根据userID获取所指导的学生信息
// @Tags Mentor
// @Accept json
// @Produce json
// @Param user_id query string false "导师的user_id"
// @Success 200 {object} common.Response
// @Router /mentor/student/mentored [get]
func ListMentoredStudents(c *gin.Context) {
	userID := c.Query("user_id")
	pagination, _ := utils.Pagination(c)
	service := mentor.ListMentoredStudentsService{}
	response := service.ListMentoredStudents(pagination, userID)
	c.JSON(200, response)
}

// @Summary 根据user_id获取导师信息
// @Tags Mentor
// @Accept json
// @Produce json
// @Param user_id query string false "导师的user_id"
// @Success 200 {object} common.Response
// @Router /mentor/info [get]
func GetMentorInfo(c *gin.Context) {
	userID := c.Query("user_id")
	service := mentor.GetMentorInfoService{}
	response := service.GetMentorInfo(userID)
	c.JSON(200, response)
}