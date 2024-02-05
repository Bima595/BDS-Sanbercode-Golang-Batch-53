// config.go

package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	// username := os.Getenv("root")
    // password := os.Getenv("bonbon123")
    // database := os.Getenv("wooo pentil iki env cok")

	username := "root"
    password := "bonbon123"
    database := "quiz3"

    connectionString := fmt.Sprintf("%s:%s@tcp(localhost:31120)/%s", username, password, database)

    db, err := sql.Open("mysql", connectionString)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Connected to the database")

	return db, nil
}
