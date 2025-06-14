package router

import (
	"geeson-auth/internal/interface/http/handler"
	"geeson-auth/internal/interface/http/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", handler.Login)

	protected := r.Group("/secure")
	protected.Use(middleware.JWTAuth())
	protected.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return r
}
