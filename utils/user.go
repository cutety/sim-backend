package utils

import "golang.org/x/crypto/bcrypt"

const (
	cost = 10
)

func ScryptPsw(password string) string {
	encodedPsw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return ""
	}
	return string(encodedPsw)
}

func DecodePsw(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}