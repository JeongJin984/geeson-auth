package config

import "os"

// GetPort returns the port to run the server on, defaulting to 8087 if not set
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8087"
	}
	return port
}

// GetJWTSecret returns the secret key for JWT token signing
// In production, this should be set as an environment variable
func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Fallback for development only - in production, this should always be set
		return "your-secret-key-for-development-only"
	}
	return secret
}
