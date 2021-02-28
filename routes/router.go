package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
		auth.PUT("user/password", v1.ChangePassword)
		auth.POST("mentor", v1.CreateMemtor)
		auth.GET("user/match/mentor", v1.GetApplyMatchingResult)
		auth.GET("user/bind/mentor", v1.ChooseMentor)
		auth.GET("info/me", v1.GetInfo)
		auth.GET("user/apply/info", v1.GetApplicationInfo)
		auth.GET("mentor/info", v1.GetMentorInfo)
		auth.GET("user/dissolve/mentor", v1.Dissolve)
		auth.PUT("student/info", v1.UpdateInfo)
		auth.GET("user/info", v1.GetUserByUserID)
		auth.GET("student/info", v1.GetStudent)
	}

	//教师权限
	mentor := auth
	mentor.Use(middlewire.JwtToken())
	mentor.Use(middlewire.AuthRole(middlewire.ROLE_TEACHER))
	{


		mentor.GET("mentor/match", v1.GetMentorMatchingResult)
		mentor.GET("mentor/student/mentored", v1.ListMentoredStudents)
		mentor.PUT("mentor/info", v1.UpdateMentor)
		mentor.GET("mentor/bind/student", v1.ChooseStudent)
	}

	// 管理员权限
	admin := r.Group("api/v1")
	admin.Use(middlewire.JwtToken())
	admin.Use(middlewire.AuthRole(middlewire.ROLE_ADMIN))
	{
		admin.GET("admin", v1.GetUserByUserID)
		admin.GET("admin/batch/mentor", v1.BatchAddMentor)
		admin.GET("admin/list/mentor", v1.ListMentors)
	}
	// 无需权限
	router := r.Group("api/v1")
	{
		router.GET("user/count", v1.InitUserPassword)
		router.POST("user/login", v1.Login)
		router.POST("student/application", v1.CreateApplication)
		router.POST("user", v1.CreateUser)
		router.POST("upload", v1.UpLoad)
	}


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	_ = r.Run(viper.GetString("server.port"))
}