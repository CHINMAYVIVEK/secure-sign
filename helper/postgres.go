package helper

import (
	"database/sql"
	"fmt"
	config "secure-sign/config"
)

// QueryRow executes a SQL query and returns a single row result
func QueryRow(query string, args ...interface{}) *sql.Row {
	return config.GetPostgreSQL().QueryRow(query, args...)
}

// Query executes a SQL query and returns the rows result
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := config.GetPostgreSQL().Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	return rows, nil
}

// Exec executes a SQL statement and returns the result
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := config.GetPostgreSQL().Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	return result, nil
}

// Insert inserts a new record into the table and returns the generated ID
func Insert(query string, args ...interface{}) (int64, error) {
	var id int64
	err := config.GetPostgreSQL().QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert record: %w", err)
	}
	return id, nil
}

// Update updates records in the table and returns the number of affected rows
func Update(query string, args ...interface{}) (int64, error) {
	result, err := config.GetPostgreSQL().Exec(query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to update records: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get affected rows: %w", err)
	}
	return rowsAffected, nil
}

// Delete deletes records from the table and returns the number of affected rows
func Delete(query string, args ...interface{}) (int64, error) {
	result, err := config.GetPostgreSQL().Exec(query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to delete records: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get affected rows: %w", err)
	}
	return rowsAffected, nil
}
