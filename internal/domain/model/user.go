package model

import (
	"errors"
	"regexp"
	"time"
)

type User struct {
	ID        int64     // Unique identifier
	Username  string    // Username
	Email     string    // Email address
	Password  string    // Password (should be stored as a hash)
	CreatedAt time.Time // Creation timestamp
}

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidEmail    = errors.New("invalid email format")
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func NewUser(id int64, username, email, password string, createdAt time.Time) (*User, error) {
	if username == "" {
		return nil, ErrInvalidUsername
	}
	if !emailRegex.MatchString(email) {
		return nil, ErrInvalidEmail
	}
	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
	}, nil
}

func (u *User) ChangeEmail(newEmail string) error {
	if !emailRegex.MatchString(newEmail) {
		return ErrInvalidEmail
	}
	u.Email = newEmail
	return nil
}
