package service

import (
	"errors"
	"geeson-auth/internal/interface/repository"
	"geeson-auth/pkg/jwt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	userRepo              repository.UserRepository
)

// SetUserRepository sets the user repository for the auth service
func SetUserRepository(repo repository.UserRepository) {
	userRepo = repo
}

func Authenticate(username, password string) (string, error) {
	if userRepo == nil {
		return "", errors.New("user repository not initialized")
	}

	user, err := userRepo.GetByUsername(username)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if user.Password != password {
		return "", ErrInvalidCredentials
	}

	return jwt.GenerateJWT(username)
}
