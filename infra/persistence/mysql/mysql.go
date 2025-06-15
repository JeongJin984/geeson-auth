package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// CreateMySqlDB establishes a connection to the MySQL database using environment variables
func CreateMySqlDB() (*sql.DB, error) {
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

	// Connection pool settings (adjust as needed)
	db.SetMaxOpenConns(25)                 // Maximum number of open connections
	db.SetMaxIdleConns(25)                 // Maximum number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime

	// Verify the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to Ping DB: %w", err)
	}

	return db, nil
}
