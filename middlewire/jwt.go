package middlewire

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"sim-backend/models/common"
	"sim-backend/utils"
	"strings"
	"time"
)

var JwtKey = []byte(viper.GetString("server.jwtKey"))

type JwtClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func SetToken(userID string) (string, common.Response) {
	expiresAt := time.Now().Add(7*24*time.Hour)
	SetClaims := JwtClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    "sim",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", utils.Response(utils.ERROR, nil)
	}
	return token, utils.Response(utils.SUCCESS, nil)
}

func CheckToken(token string) (*JwtClaims, int) {
	var claims JwtClaims
	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, utils.ERROR_TOKEN_WRONG
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, utils.ERROR_TOKEN_RUNTIME
			} else {
				return nil, utils.ERROR_TOKEN_TYPE_WRONG
			}
		}
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*JwtClaims); ok && setToken.Valid {
			return key, utils.SUCCESS
		} else {
			return nil, utils.ERROR_TOKEN_WRONG
		}
	}

	return nil, utils.ERROR_TOKEN_WRONG
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			c.JSON(200, utils.Response(utils.ERROR_TOKEN_EXIST, nil))
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")
		if len(checkToken) == 0 {
			c.JSON(200, utils.Response(utils.ERROR_TOKEN_TYPE_WRONG, nil))
			c.Abort()
			return
		}

		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			c.JSON(200, utils.Response(utils.ERROR_TOKEN_TYPE_WRONG, nil))
			c.Abort()
			return
		}

		key, code := CheckToken(checkToken[1])
		if code != utils.SUCCESS {
			c.JSON(200, utils.Response(code, nil))
			c.Abort()
			return
		}

		c.Set("user_id", key)
		c.Next()
	}
}