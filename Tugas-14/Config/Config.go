package Config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "Bima123yayang"
	database = "student"
)

var dsn = fmt.Sprintf("%v:%v@/%v", username, password, database)

// MySQL creates a MySQL database connection
func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Connected to the database")
	return db, nil
}
