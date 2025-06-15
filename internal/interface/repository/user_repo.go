package repository

import "geeson-auth/internal/domain/model"

// UserRepository is a port (interface) for user data access
type UserRepository interface {
	// GetByID retrieves a user by their ID
	GetByID(id int64) (*model.User, error)

	// GetByUsername retrieves a user by their username
	GetByUsername(username string) (*model.User, error)
}
