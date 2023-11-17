package mysql

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewDBPool(username, password, host, port, database string, maxOpenConns, maxIdleConns int) (*DBPool, error) {
	// Build MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	// Initialize MySQL driver
	// Note: The driver is only imported with a blank identifier ("_")
	// as it is used for its side effects (registering with sql package).
	_, _ = sql.Open("mysql", dsn)

	// Open a connection pool
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("database Ping sucessfull")
	return &DBPool{db}, nil
}
