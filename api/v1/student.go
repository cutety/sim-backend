package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/student"
)

func CreateApplication(c *gin.Context) {
	var service student.CreateApplicationService
	if err := c.ShouldBindJSON(&service); err == nil {
		response := service.CreateApplication()
		c.JSON(200, response)
	} else {
		c.JSON(200, common.Response{Error: err.Error()})
	}
}
