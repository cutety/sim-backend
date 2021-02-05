package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/user"
	"sim-backend/utils"
	"sim-backend/utils/validator"
)

func GetUserByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	service := user.GetUserByUserIDService{}
	response := service.GetUserByUserIDService(userID)
	c.JSON(200, response)
}

func GetUserByID(c *gin.Context) {

}

func ChangePassword(c *gin.Context) {
	service := &user.ChangePasswordService{}
	if err := c.ShouldBindJSON(service); err == nil {
		response := service.UserChangePasswordByUserID()
		c.JSON(200, response)
	} else {
		c.JSON(200, common.Response{Error: err.Error()})
	}
}

func InitUserPassword(c *gin.Context) {
	service := user.InitUserPswService{}
	err := service.InitUserPsw()
	if err != nil {
		c.JSON(200, common.Response{Error: err.Error()})
	} else {
		c.JSON(200, common.Response{Msg: "suc"})
	}
}

func Login(c *gin.Context) {
	var service user.LoginService
	_ = c.ShouldBindJSON(&service)
	response := service.Login()
	c.JSON(200, response)
}


func CreateUser(c *gin.Context) {
	var service user.CreateUserService
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
		response := service.CreateUser()
		c.JSON(200, response)
	} else {
		c.JSON(200, utils.Response(utils.ERROR, err))
	}
}

func GetApplyMatchingResult(c *gin.Context) {
	userID := c.Param("user_id")
	service := user.GetApplyMatchingResultService{}
	response := service.GetApplyMatchingResult(userID)
	c.JSON(200, response)
}

func ChooseMentor(c *gin.Context) {
	userID := c.Query("user_id")
	mentorUserID := c.Query("mentor_user_id")
	service := user.ChooseMentorService{}
	response := service.ChooseMentor(userID, mentorUserID)
	c.JSON(200, response)
}