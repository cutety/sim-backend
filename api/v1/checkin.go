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
// @Param CheckinInfo body checkin.NewStudentCheckinService true "报道信息"
// @Success 200 {object} common.Response
// @Router /students/checkin/new [post]
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

// @Summary 男生女生数量
// @Tags Checkin
// @version 1.0
// @Accept application/x-json-stream
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/gender/amount/{grade} [get]
func GetMaleAndFemaleAmount(c *gin.Context) {
	service := &checkin.GetMaleAndFemaleAmountService{}
	grade := c.Param("grade")
	amount, err := service.GetMaleAndFemaleAmount(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}


// @Summary 根据年级获取学生人数
// @Tags Checkin
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/amount/{grade} [get]
func GetStudentsAmountByGrade(c *gin.Context) {
	service := &checkin.GetStudentsAmountByGradeService{}
	grade := c.Param("grade")
	amount, err := service.GetStudentsAmountByGrade(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}

// @Summary 根据年级获取年龄分布
// @Tags Checkin
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/age/distribution/{grade} [get]
func GetAgeDistribution(c *gin.Context) {
	service := &checkin.GetAgeDistributionService{}
	grade := c.Param("grade")
	amount, err := service.GetAgeDistribution(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}

// @Summary 根据年级获取学生信息轮播表
// @Tags Checkin
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/info/table/{grade} [get]
func GetStudentsInfoTable(c *gin.Context) {
	service := &checkin.GetStudentsInfoTableService{}
	grade := c.Param("grade")
	amount, err := service.GetStudentsInfoTable(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}

// @Summary 根据年级获取省份信息
// @Tags Checkin
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/province/{grade} [get]
func GetStudentsProvince(c *gin.Context) {
	service := &checkin.GetStudentsProvinceService{}
	grade := c.Param("grade")
	amount, err := service.GetStudentsProvince(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}

// @Summary 根据年级获取姓排行
// @Tags Checkin
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/firstname/{grade} [get]
func GetFirstnameByGrade(c *gin.Context) {
	service := &checkin.GetFirstNameByGradeService{}
	grade := c.Param("grade")
	amount, err := service.GetFirstNameByGrade(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}

// @Summary 根据年级获取同名情况
// @Tags Checkin
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/same/name/{grade} [get]
func GetSameNameByGrade(c *gin.Context) {
	service := &checkin.GetSameNameByGradeService{}
	grade := c.Param("grade")
	amount, err := service.GetSameNameByGrade(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}

// @Summary 根据年级获取同名情况
// @Tags Checkin
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/same/birthday/{grade} [get]
func GetSameBirthdayByGrade(c *gin.Context) {
	service := &checkin.GetSameBirthdayByGradeService{}
	grade := c.Param("grade")
	amount, err := service.GetSameBirthdayByGrade(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}

// @Summary 根据年级获取专业排行
// @Tags Checkin
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/major/rank/{grade} [get]
func GetMajorRankByGrade(c *gin.Context) {
	service := &checkin.GetMajorRankByGradeService{}
	grade := c.Param("grade")
	amount, err := service.GetMajorRankByGrade(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}

// @Summary 根据年级获取报道信息
// @Tags Checkin
// @Param grade path string true "年级"
// @Success 200 object common.Response 成功后返回值
// @Router /students/checkin/info/{grade} [get]
func GetCheckinInfoByGrade(c *gin.Context) {
	service := &checkin.GetCheckinInfoByGradeService{}
	grade := c.Param("grade")
	amount, err := service.GetCheckinInfoByGrade(grade)
	if err != nil {
		c.JSON(200, utils.ResponseWithError(utils.ERROR, err))
		c.Abort()
		return
	}
	c.JSON(200, utils.Response(utils.SUCCESS, amount))
}