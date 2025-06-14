package jwt

import (
	"geeson-auth/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetJWTSecret()))
}
