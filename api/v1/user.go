package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/user"
)

func GetUserByUserID(c *gin.Context) {

}

func GetUserByID(c *gin.Context) {

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
	var service user.UserLoginService
	_ = c.ShouldBindJSON(&service)
	response := service.Login()
	c.JSON(200, response)
}
