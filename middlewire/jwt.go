package middlewire

import (
	"errors"
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
	Role int `json:"role"`
	jwt.StandardClaims
}

func SetToken(userID string, role int) (string, common.Response) {
	expiresAt := time.Now().Add(7*24*time.Hour)
	SetClaims := JwtClaims{
		userID,
		role,
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

func GetTokenFromRequestHeader(c *gin.Context) string {
	tokenHeader := c.Request.Header.Get("Authorization")
	if tokenHeader == "" {
		c.JSON(200, utils.Response(utils.ERROR_TOKEN_EXIST, nil))
		c.Abort()
		return ""
	}

	token := strings.Split(tokenHeader, " ")
	if len(token) == 0 {
		c.JSON(200, utils.Response(utils.ERROR_TOKEN_TYPE_WRONG, nil))
		c.Abort()
		return ""
	}

	if len(token) != 2 && token[0] != "Bearer" {
		c.JSON(200, utils.Response(utils.ERROR_TOKEN_TYPE_WRONG, nil))
		c.Abort()
		return ""
	}
	return token[1]
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetTokenFromRequestHeader(c)
		if token != "" {
			key, code := CheckToken(token)
			if code != utils.SUCCESS {
				c.JSON(200, utils.Response(code, nil))
				c.Abort()
				return
			}
			c.Set("user_info", key)
			c.Set("user_id", key.UserID)
			c.Next()
		}
	}
}

func ParseJwtToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("unexpected token claims")
		}
		return JwtKey, nil
	})

	return token, err
}

// GetValueFromToken get value form jwt token string
func GetValueFromToken(tokenString, key string) (value interface{}, found bool) {
	jwtToken, err := ParseJwtToken(tokenString)
	if err != nil {
		return
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok {
		if v, ok := claims[key]; ok {
			return v, true
		}
	}

	return
}