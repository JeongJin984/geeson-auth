package repository

import "geeson-auth/internal/domain/model"

// UserRepository 는 user 조회용 포트(인터페이스)입니다.
type UserRepository interface {
	GetByID(id int64) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
}
