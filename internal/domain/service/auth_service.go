package service

import (
	"errors"
	"geeson-auth/pkg/jwt"
)

var dummyUsers = map[string]string{
	"user1": "pass1",
	"admin": "admin123",
}

func Authenticate(username, password string) (string, error) {
	if pw, ok := dummyUsers[username]; ok && pw == password {
		return jwt.GenerateJWT(username)
	}
	return "", errors.New("invalid credentials")
}
