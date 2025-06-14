package handler

import (
	"geeson-auth/internal/domain/model"
	"geeson-auth/internal/domain/service"
	"geeson-auth/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.L().Warn("Invalid login payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	token, err := service.Authenticate(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
