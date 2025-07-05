package usecase

import (
	"errors"
	"geeson-auth/internal/domain/repository"
	"geeson-auth/internal/domain/service"
	"geeson-auth/pkg/jwt"
)

type AuthUseCase struct {
	userRepo    repository.UserRepository
	authService *service.AuthService
}

func NewAuthUseCase(userRepo repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo:    userRepo,
		authService: service.NewAuthService(),
	}
}

func (uc *AuthUseCase) Authenticate(username, password string) (string, error) {
	if uc.userRepo == nil {
		return "", errors.New("user repository not initialized")
	}

	user, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return "", service.ErrInvalidCredentials
	}

	if !uc.authService.CheckPasswordHash(password, user.Password) {
		return "", service.ErrInvalidCredentials
	}

	return jwt.GenerateJWT(username)
}
