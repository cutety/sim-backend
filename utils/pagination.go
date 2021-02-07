package utils

import (
	"github.com/gin-gonic/gin"
	"sim-backend/models/common"
	"strconv"
)

func Pagination(c *gin.Context) (*common.Pagination,error) {
	limit := c.DefaultQuery("page_size", "8")
	pageNumber := c.DefaultQuery("page_number", "1")
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 0 {
		return nil,err
	}
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil || pageNumberInt < 0 {
		return nil,err
	}
	order := c.DefaultQuery("order", "desc")
	return &common.Pagination{
		Page:  pageNumberInt,
		Limit: limitInt,
		Order: order,
	}, err
}
