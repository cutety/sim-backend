package middlewire

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"sim-backend/utils"
	"sim-backend/utils/logger"
)

const (
	ROLE_ADMIN   = 1
	ROLE_STUDENT = 2
	ROLE_TEACHER = 3
)

var tokenRole = [][]int{{1,2,3}, {2}, {2,3}}

func AuthRole(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetTokenFromRequestHeader(c)
		r, found := GetValueFromToken(token, "role")
		if found {
			tr := cast.ToInt(r)
			logger.Info("role request:", tr)
			logger.Info("role require:", role)
			logger.Info("role range:", tokenRole[tr-1])
			for index, k := range tokenRole[tr-1] {
				if k == role {
					c.Set("role", k)
					c.Next()
					break
				}
				if index == (len(tokenRole[tr-1]) -1) {
					c.JSON(200, utils.Response(utils.ERROR_USER_AUTHORITY, nil))
					c.Abort()
					return
				}
			}
		} else {
			c.JSON(200, utils.Response(utils.ERROR_TOKEN_WRONG, nil))
			c.Abort()
			return
		}
	}
}
