package v1

import (
	"github.com/gin-gonic/gin"
	"sim-backend/service/evaluation"
	"sim-backend/utils"
)

// CreateEvaluation 创建评价
// @Summary 开始上课
// @Tags Evaluation
// @Accept json
// @Produce json
// @Param CreateEvaluationService body evaluation.CreateEvaluationService true "评价信息"
// @Success 200 {object} common.Response
// @Router /student/add/evaluation [POST]
func CreateEvaluation(c *gin.Context) {
	service := &evaluation.CreateEvaluationService{}
	if err := c.ShouldBindJSON(service); err == nil {
		err := service.CreateEvaluation()
		c.JSON(200, utils.ResponseAll(nil, err))
	} else {
		c.JSON(200, utils.ResponseAll(nil, err))
	}
}
