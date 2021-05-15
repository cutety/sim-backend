package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/mentor"
	"sim-backend/utils"
	"sim-backend/utils/validator"
	"strconv"
)

// @Summary 添加导师
// @Tags Mentor
// @Accept json
// @Produce json
// @Param mentor body mentor.CreateMentorService true "导师的个人信息"
// @Success 200 {object} common.Response
// @Router /mentor [post]
func CreateMentor(c *gin.Context) {
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
// @Router /mentor/info [put]
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
	response := service.ListMentoredStudents(pagination, userID, common.IS_MATCHED)
	c.JSON(200, response)
}

// @Summary 根据userID获取指导申请
// @Tags Mentor
// @Accept json
// @Produce json
// @Param user_id query string false "导师的user_id"
// @Success 200 {object} common.Response
// @Router /mentor/student/request [get]
func ListInstructRequest(c *gin.Context) {
	userID := c.Query("user_id")
	pagination, _ := utils.Pagination(c)
	service := mentor.ListMentoredStudentsService{}
	response := service.ListMentoredStudents(pagination, userID, common.NOT_MATCHED)
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

// @Summary 老师选学生
// @Tags Mentor
// @Accept json
// @Produce json
// @Param user_id query string true "学生user_id"
// @Param mentor_user_id query string true "导师user_id"
// @Success 200 {object} common.Response
// @Router /mentor/bind/student [get]
func ChooseStudent(c *gin.Context) {
	userID := c.Query("user_id")
	mentorUserID := c.Query("mentor_user_id")
	status,_ := strconv.Atoi(c.Query("status"))

	service := mentor.ChooseStudentService{}
	response := service.ChooseStudent(userID, mentorUserID, status)
	c.JSON(200, response)
}


// DeleteMentorByID 删除导师
func DeleteMentorByID(c *gin.Context) {
	userID := c.Query("mentor_id")
	if userID == "" {
		c.JSON(200, utils.ResponseAll(nil , errors.New(utils.GetErrMsg(utils.ERROR_DELETE_ERROR_USER_ID_EXIST))))
		c.Abort()
		return
	}
	service := mentor.DeleteMentorByIDService{}
	err := service.DeleteMentorByID(userID)
	c.JSON(200, utils.ResponseAll(nil, err))
}