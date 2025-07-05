package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"geeson-auth/internal/domain/repository"
	"geeson-auth/pkg/logger"

	"geeson-auth/internal/domain/model"
)

var _ repository.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

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
