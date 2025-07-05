package router

import (
	"geeson-auth/internal/application"
	"geeson-auth/internal/domain/repository"
	"geeson-auth/internal/presentation/http/handler"
	"geeson-auth/internal/presentation/http/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures the HTTP routes and returns the router
func SetupRouter(userRepo repository.UserRepository) *gin.Engine {
	r := gin.Default()

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(userRepo)

	// Initialize controllers
	authController := handler.NewAuthController(authUseCase)

	// Public routes
	r.POST("/login", authController.Login)

	// Protected routes
	protected := r.Group("/secure")
	protected.Use(middleware.JWTAuth())
	protected.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return r
}
