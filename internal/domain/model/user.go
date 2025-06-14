package model

import (
	"errors"
	"regexp"
	"time"
)

// User 도메인 객체는 사용자에 대한 핵심 속성과 비즈니스 규칙을 캡슐화합니다.
type User struct {
	ID        int64     // 고유 식별자
	Username  string    // 사용자 이름
	Email     string    // 이메일 주소
	Password  string    // 비밀번호
	CreatedAt time.Time // 생성 시각
}

var (
	// 도메인 오류 정의
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidEmail    = errors.New("invalid email format")
)

// 이메일 유효성 검사를 위한 정규식 (단순 검증 예시)
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// NewUser 는 User 도메인 객체를 생성하면서 기본 검증을 수행합니다.
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

// ChangeEmail 은 이메일을 변경하며, 도메인 규칙을 검증합니다.
func (u *User) ChangeEmail(newEmail string) error {
	if !emailRegex.MatchString(newEmail) {
		return ErrInvalidEmail
	}
	u.Email = newEmail
	return nil
}
