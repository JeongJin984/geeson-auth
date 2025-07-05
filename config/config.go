package config

import "os"

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8087"
	}
	return port
}

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "your-secret-key-for-development-only"
	}
	return secret
}
