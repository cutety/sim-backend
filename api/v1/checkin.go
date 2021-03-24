package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"sim-backend/service/checkin"
	"sim-backend/utils"
	"sim-backend/utils/validator"
)

// @Summary 新生报到
// @Tags Checkin
// @Accept json
// @Produce json
// @Param CheckinInfo body checkin.CheckinService true "报道信息"
// @Success 200 {object} common.Response
// @Router /checkin/new [post]
func StudentCheckin(c *gin.Context) {
	service := &checkin.NewStudentCheckinService{}
	if err := c.ShouldBindJSON(service); err == nil {
		msg, code := validator.Validate(service)
		if code != utils.SUCCESS {
			c.JSON(200, common.Response{
				Status: code,
				Msg:    msg,
			})
			c.Abort()
			return
		}
		err := service.CheckIn()
		if err != nil {
			c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		}
		c.JSON(200, common.Response{
			Status: utils.SUCCESS,
			Data:   nil,
			Msg:    "签到成功",
			Error:  "",
		})
	} else {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
}

func GetCheckinAmount(c *gin.Context) {
	service := &checkin.GetCheckinAmountService{}
	grade := c.Query("grade")
	total, err := service.GetCheckinAmount(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, common.Response{
		Status: utils.ERROR,
		Data:   total,
		Msg:    "",
		Error:  "",
	})
}