package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"sim-backend/models/common"
	"sim-backend/service/user"
	"sim-backend/utils"
	"sim-backend/utils/logger"
	"sim-backend/utils/validator"
)

// @Summary 根据user_id获取用户信息
// @Tags User
// @Accept json
// @Produce json
// @Param user_id query string true "user_id"
// @Success 200 {object} common.Response
// @Router /user/info/:user_id [get]
func GetUserByUserID(c *gin.Context) {
	userID := c.Query("user_id")
	service := user.GetUserByUserIDService{}
	response := service.GetUserByUserIDService(userID)
	c.JSON(200, response)
}

// @Summary 修改密码
// @Tags User
// @Accept json
// @Produce json
// @Param ChangePassword body user.ChangePasswordService true "修改密码RequestBody"
// @Success 200 {object} common.Response
// @Router /user/password [post]
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

// @Summary 用户登录
// @Tags User
// @Accept json
// @Produce json
// @Param LoginService body user.LoginService true "登录参数"
// @Success 200 {object} common.Response
// @Router /user/login [post]
func Login(c *gin.Context) {
	var service user.LoginService
	_ = c.ShouldBindJSON(&service)
	response := service.Login()
	c.JSON(200, response)
}

// @Summary 创建用户
// @Tags User
// @Accept json
// @Produce json
// @Param CreateUserService body user.CreateUserService true "创建用户参数"
// @Success 200 {object} common.Response
// @Router /user [post]
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



// @Summary 解除关系
// @Tags User
// @Accept json
// @Produce json
// @Param user_id query string true "学生user_id"
// @Success 200 {object} common.Response
// @Router /user/dissolve/mentor [get]
func Dissolve(c *gin.Context) {
	userID := c.Query("user_id")
	service := user.DissolveService{}
	response := service.Dissolve(userID)
	c.JSON(200, response)
}

// @Summary 根据token获取用户信息
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Router /info/me [get]
func GetInfo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	logger.Info("userId", userID)
	service := user.GetUserByUserIDService{}
	response := service.GetUserByUserIDService(cast.ToString(userID))
	c.JSON(200, response)
}

// @Summary 师生双选
// @Tags User
// @Accept json
// @Produce json
// @Param application body user.UpdateApplicationService true "师生双选Request"
// @Success 200 {object} common.Response
// @Router /user/dual/select [put]
func DualSelect(c *gin.Context) {
	service := &user.UpdateApplicationService{}
	if err := c.ShouldBindJSON(service); err == nil {
		response := service.UpdateApplication()
		c.JSON(200, response)
	} else {
		c.JSON(200, common.Response{Error: err.Error()})
	}
}