package database

import (
	"database/sql"
	"time"

	/*
	 required for DB.
	*/
	_ "github.com/go-sql-driver/mysql"
)

// Student struct.
type Student struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//InitDB initializes the DB connection.
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1234@/students")
	if err != nil {
		return nil, err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
