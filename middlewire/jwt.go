package middlewire

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

var jwtKey = []byte(viper.GetString("server.jwtKey"))


type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}