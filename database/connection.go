package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	fmt.Println("✅ Successfully connected to database!")
	return db, nil
}

func CreateTables(db *sql.DB) error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		age INT NOT NULL,
		phone VARCHAR(20),
		is_hidden BOOLEAN DEFAULT FALSE,
		rating DECIMAL(3, 2),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(createUsersTable)
	if err != nil {
		return fmt.Errorf("error creating users table: %w", err)
	}

	fmt.Println("✅ Tables are ready!")
	return nil
}
