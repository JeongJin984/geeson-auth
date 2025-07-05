package main

import (
	"database/sql"
	"geeson-auth/config"
	"geeson-auth/infra/persistence/mysql"
	"geeson-auth/infra/router"
	"geeson-auth/pkg/logger"
)

func main() {
	logger.InitLogger()

	db, err := mysql.CreateMySqlDB()
	if err != nil {
		logger.L().Error("Failed to connect to database: " + err.Error())
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.L().Error("Failed to close database connection: " + err.Error())
		}
	}(db)

	userRepo := mysql.NewUserRepository(db)

	r := router.SetupRouter(userRepo)
	port := config.GetPort()
	logger.L().Info("Starting server on port " + port)
	err = r.Run(":" + port)
	if err != nil {
		logger.L().Error("Error starting server: " + err.Error())
	}
}
