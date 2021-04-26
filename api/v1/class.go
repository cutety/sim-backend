package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/service/class"
	"sim-backend/utils"
)

func CreateClass(c *gin.Context) {
	service := &class.CreateClassService{}
	if err := c.ShouldBindJSON(service); err == nil {
		err := service.CreateClass()
		c.JSON(200, utils.ResponseAll(nil, err))
	} else {
		c.JSON(200, utils.ResponseAll(nil ,err))
	}
}

// ListClassByGrade 通过年级获取班级列表
func ListClassByGrade(c *gin.Context) {
	grade := c.Query("grade")
	service := &class.ListClassesByGradeService{}
	result, err := service.ListClassesByGrade(grade)
	c.JSON(200, utils.ResponseAll(result, err))
}
