package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/middlewire"
	"sim-backend/models/common"
	"sim-backend/service/user"
	"sim-backend/utils"
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
	var token string
	var res common.Response
	if response.Status == utils.SUCCESS {
		token, res = middlewire.SetToken(service.UserID)
	}
	c.JSON(200, common.LoginResponse{
		Status: res.Status,
		Msg:    res.Msg,
		Error:  res.Msg,
		Token:  token,
	})
}
