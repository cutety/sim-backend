package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	v1 "sim-backend/api/v1"
	"sim-backend/utils/logger"
)

func InitRouter() {
	gin.SetMode(viper.GetString("server.mode"))
	r := gin.New()
	r.Use(logger.GinLogger())

	router := r.Group("api/v1")
	{
		router.GET("user/count", v1.InitUserPassword)
		router.POST("user/login", v1.Login)
	}

	_ = r.Run(viper.GetString("server.port"))
}