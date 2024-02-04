package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var pgDB *sql.DB

// ConnectPostgreSQL initializes and configures the PostgreSQL database connection.
func ConnectPostgreSQL() error {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Cfg.Database.Host, Cfg.Database.Port, Cfg.Database.Username, Cfg.Database.Password, Cfg.Database.Name)

	dbobj, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	// Set connection pool properties
	dbobj.SetMaxOpenConns(100)
	dbobj.SetMaxIdleConns(20)
	dbobj.SetConnMaxLifetime(5 * time.Minute)

	// Check database connection
	if err := dbobj.Ping(); err != nil {
		dbobj.Close()
		return fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	pgDB = dbobj
	return nil
}

// GetPostgreSQL returns the PostgreSQL database instance.
func GetPostgreSQL() *sql.DB {
	return pgDB
}

// ClosePostgreSQL closes the PostgreSQL database connection.
func ClosePostgreSQL() {
	if pgDB != nil {
		pgDB.Close()
	}
}
