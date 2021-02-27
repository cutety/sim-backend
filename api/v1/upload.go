package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/service/admin"
)

func UpLoad(c *gin.Context) {
	file, _, _ := c.Request.FormFile("file")
	service := admin.BatchAddMentorService{}
	response := service.BatchAddMentor(file)
	c.JSON(200, response)
}