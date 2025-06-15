package handler

import (
	"geeson-auth/internal/domain/model"
	"geeson-auth/internal/usercase"
	"geeson-auth/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthController handles HTTP requests related to authentication
type AuthController struct {
	authUseCase *usercase.AuthUseCase
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authUseCase *usercase.AuthUseCase) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}

// Login handles user login requests
func (c *AuthController) Login(ctx *gin.Context) {
	var req model.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.L().Warn("Invalid login payload")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	token, err := c.authUseCase.Authenticate(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
