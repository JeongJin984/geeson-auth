package config

import "os"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func GetJWTSecret() string {
	return "your-secret-key"
}
