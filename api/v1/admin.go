package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/service/admin"
	"sim-backend/utils"
)

// @Summary 批量添加导师
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id query string true "学生user_id"
// @Param mentor_user_id query string true "导师user_id"
// @Success 200 {object} common.Response
// @Router /admin/batch/mentor [get]
func BatchAddMentor(c *gin.Context) {
	//service := admin.BatchAddMentorService{}
	//response := service.BatchAddMentor()
	//c.JSON(200, response)
}

func ListMentors(c *gin.Context) {
	pagination, _ := utils.Pagination(c)
	service := admin.ListMentorsService{}
	response := service.ListMentors(pagination)
	c.JSON(200, response)
}