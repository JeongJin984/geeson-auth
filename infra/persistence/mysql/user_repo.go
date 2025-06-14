package mysql

import (
	"database/sql"
	"fmt"

	"geeson-auth/internal/domain/model"
	"geeson-auth/internal/interface/repository"
)

// Ensure UserRepository implements repository.UserRepository
var _ repository.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *sql.DB
}

// NewUserRepository 는 포트 인터페이스 타입으로 반환합니다.
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int64) (*model.User, error) {
	const q = `
        SELECT id, username, email, created_at
          FROM users
         WHERE id = ?`
	row := r.db.QueryRow(q, id)

	u := &model.User{}
	if err := row.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found (id=%d)", id)
		}
		return nil, fmt.Errorf("row.Scan failed: %w", err)
	}
	return u, nil
}
