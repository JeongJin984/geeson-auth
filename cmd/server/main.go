package main

import (
	"geeson-auth/config"
	"geeson-auth/infra/router"
	"geeson-auth/pkg/logger"
)

func main() {
	logger.InitLogger()
	r := router.SetupRouter()
	port := config.GetPort()
	logger.L().Info("Starting server on port " + port)
	err := r.Run(":" + port)
	if err != nil {
		logger.L().Info("Error starting server")
	}
}
