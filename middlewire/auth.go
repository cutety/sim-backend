package middlewire

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"sim-backend/utils"
)

const (
	ROLE_ADMIN   = 1
	ROLE_STUDENT = 2
	ROLE_TEACHER = 3
)

var tokenRole = [][]int{{1, 2, 3}, {1}, {1, 2}}

func AuthRole(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetTokenFromRequestHeader(c)
		r, found := GetValueFromToken(token, "role")
		if found {
			tr := cast.ToInt(r)
			for _, k := range tokenRole[tr-1] {
				if k == role {
					c.Set("role", k)
					c.Next()
					break
				}
			}
		} else {
			c.JSON(200, utils.Response(utils.ERROR_USER_AUTHORITY, nil))
			c.Abort()
			return
		}
	}
}
