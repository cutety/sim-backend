package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/student"
	"sim-backend/service/user"
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

// @Summary 获取学生列表
// @Tags Student
// @Accept json
// @Produce json
// @Param pagination.page query string false "page"
// @Param pagination.limit query string false "limit"
// @Param stu_name query string false "学生姓名"
// @Param grade query string false "年级"
// @Param major query string false "专业"
// @Success 200 {object} common.Response
// @Router /student/detail [get]
func ListStudentsDetail(c *gin.Context) {
	pagination, _ := utils.Pagination(c)
	service := student.ListStudentsDetailService{}
	if err := c.BindQuery(&service); err == nil {
		apps, total, err := service.ListStudentsDetail(pagination)
		if err != nil {
			c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
			return
		}
		list := common.DataList{
			Items: apps,
			Total: total,
		}
		c.JSON(200, utils.Response(utils.SUCCESS, list))
	} else {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
	}
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

// @Summary 学生选老师
// @Tags Student
// @Accept json
// @Produce json
// @Param user_id query string true "学生user_id"
// @Param mentor_user_id query string true "导师user_id"
// @Success 200 {object} common.Response
// @Router /student/bind/mentor [get]
func ChooseMentor(c *gin.Context) {
	userID := c.Query("user_id")
	mentorUserID := c.Query("mentor_user_id")
	service := user.ChooseMentorService{}
	response := service.ChooseMentor(userID, mentorUserID)
	c.JSON(200, response)
}

// @Summary 查看历史录取记录
// @Tags Student
// @Accept json
// @Produce json
// @Param user_id query string true "学生user_id"
// @Success 200 {object} common.Response
// @Router /student/admission/history [get]
func ListMatchedAdmittedStudents(c *gin.Context) {
	pagination, _ := utils.Pagination(c)
	userID := c.Query("user_id")
	service := student.ListMatchedAdmittedStudentsService{}
	students, total, err := service.ListMatchedAdmittedStudents(userID, pagination)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		return
	}
	list := common.DataList{
		Items: students,
		Total: total,
	}
	c.JSON(200, utils.Response(utils.SUCCESS, list))
}

// @Summary 通过年级获取班级列表
// @Tags User
// @Accept json
// @Produce json
// @Param grade query string true "班级"
// @Success 200 {object} common.Response
// @Router /user/list/class [get]
func ListClassesByGrade(c *gin.Context) {
	grade := c.Query("grade")
	service := student.ListClassesByGradeService{}
	classes, err := service.ListClassesByGrade(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, classes))
}