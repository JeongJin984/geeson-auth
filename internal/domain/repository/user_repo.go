package repository

import "geeson-auth/internal/domain/model"

type UserRepository interface {
	GetByID(id int64) (*model.User, error)

	GetByUsername(username string) (*model.User, error)
}
