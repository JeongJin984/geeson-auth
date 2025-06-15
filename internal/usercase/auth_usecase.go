package usercase

import (
	"errors"
	"geeson-auth/internal/domain/service"
	"geeson-auth/internal/interface/repository"
	"geeson-auth/pkg/jwt"
)

// AuthUseCase encapsulates the authentication use cases
type AuthUseCase struct {
	userRepo    repository.UserRepository
	authService *service.AuthService
}

// NewAuthUseCase creates a new instance of AuthUseCase
func NewAuthUseCase(userRepo repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo:    userRepo,
		authService: service.NewAuthService(),
	}
}

// Authenticate validates user credentials and returns a JWT token if valid
func (uc *AuthUseCase) Authenticate(username, password string) (string, error) {
	if uc.userRepo == nil {
		return "", errors.New("user repository not initialized")
	}

	user, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return "", service.ErrInvalidCredentials
	}

	// Check if the provided password matches the stored hash
	if !uc.authService.CheckPasswordHash(password, user.Password) {
		return "", service.ErrInvalidCredentials
	}

	return jwt.GenerateJWT(username)
}
