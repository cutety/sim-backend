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

	auth := r.Group("api/v1")
	auth.Use(middlewire.JwtToken())
	{
		auth.GET("user/info/:user_id", v1.GetUserByUserID)
		auth.POST("user/password", v1.ChangePassword)
	}

	router := r.Group("api/v1")
	{
		router.GET("user/count", v1.InitUserPassword)
		router.POST("user/login", v1.Login)

	}

	_ = r.Run(viper.GetString("server.port"))
}