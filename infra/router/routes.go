package router

import (
	"geeson-auth/internal/interface/http/handler"
	"geeson-auth/internal/interface/http/middleware"
	"geeson-auth/internal/interface/repository"
	"geeson-auth/internal/usercase"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures the HTTP routes and returns the router
func SetupRouter(userRepo repository.UserRepository) *gin.Engine {
	r := gin.Default()

	// Initialize use cases
	authUseCase := usercase.NewAuthUseCase(userRepo)

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
