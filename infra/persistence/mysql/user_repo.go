package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"geeson-auth/pkg/logger"

	"geeson-auth/internal/domain/model"
	"geeson-auth/internal/interface/repository"
)

// Ensure UserRepository implements repository.UserRepository
var _ repository.UserRepository = (*UserRepository)(nil)

// UserRepository implements the repository.UserRepository interface using MySQL
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository instance that implements the port interface
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

// GetByID retrieves a user by their ID
func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	const q = `
        SELECT id, username, email, password_hash, created_at
          FROM users
         WHERE id = ?`
	row := r.db.QueryRow(q, id)

	u := &model.User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.L().Error(fmt.Sprintf("user not found (id=%d)", id))
			return nil, fmt.Errorf("user not found (id=%d)", id)
		}
		return nil, fmt.Errorf("row.Scan failed: %w", err)
	}
	return u, nil
}

// GetByUsername retrieves a user by their username
func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	const q = `
        SELECT id, username, email, password_hash, created_at
          FROM users
         WHERE username = ?`
	row := r.db.QueryRow(q, username)

	u := &model.User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.L().Error(fmt.Sprintf("user not found (username=%s)", username))
			return nil, fmt.Errorf("user not found (username=%s)", username)
		}
		return nil, fmt.Errorf("row.Scan failed: %w", err)
	}
	return u, nil
}
