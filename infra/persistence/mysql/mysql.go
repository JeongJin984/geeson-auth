package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySqlDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")     // e.g. "root"
	password := os.Getenv("DB_PASS") // e.g. "secret"
	host := os.Getenv("DB_HOST")     // e.g. "127.0.0.1"
	port := os.Getenv("DB_PORT")     // e.g. "3306"
	dbname := os.Getenv("DB_NAME")   // e.g. "myapp"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		user, password, host, port, dbname,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to Open DB: %w", err)
	}

	// 커넥션 풀 설정 (필요에 맞게 조정)
	db.SetMaxOpenConns(25)                 // 최대 열려있는 커넥션 수
	db.SetMaxIdleConns(25)                 // 최대 대기 커넥션 수
	db.SetConnMaxLifetime(5 * time.Minute) // 커넥션 최대 수명

	// 실제 연결 확인
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to Ping DB: %w", err)
	}

	return db, nil
}
