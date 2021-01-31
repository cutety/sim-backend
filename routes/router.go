package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	v1 "sim-backend/api/v1"
	"sim-backend/middlewire"
	"sim-backend/utils/logger"
)

func InitRouter() {
	gin.SetMode(viper.GetString("server.mode"))
	r := gin.New()
	r.Use(logger.GinLogger())
	r.Use(gin.Recovery())
	r.Use(middlewire.Cors())

	// 需要认证的
	auth := r.Group("api/v1")
	auth.Use(middlewire.JwtToken())
	{
		//auth.GET("user/info/:user_id", v1.GetUserByUserID)
		auth.POST("user/password", v1.ChangePassword)
		auth.POST("mentor", v1.CreateMemtor)
	}

	//教师权限
	teacher := auth
	teacher.Use(middlewire.JwtToken())
	teacher.Use(middlewire.AuthRole(middlewire.ROLE_TEACHER))
	{
		teacher.GET("user/info/:user_id", v1.GetUserByUserID)
	}

	// 管理员权限
	admin := r.Group("api/v1")
	admin.Use(middlewire.JwtToken())
	admin.Use(middlewire.AuthRole(middlewire.ROLE_ADMIN))
	{
		admin.GET("admin", v1.GetUserByUserID)
	}

	// 无需权限
	router := r.Group("api/v1")
	{
		router.GET("user/count", v1.InitUserPassword)
		router.POST("user/login", v1.Login)
		router.POST("student/application", v1.CreateApplication)
	}

	_ = r.Run(viper.GetString("server.port"))
}