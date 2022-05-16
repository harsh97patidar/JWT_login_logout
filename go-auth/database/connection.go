package database

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"
)

// DB is a global variable to hold db connection
var DB *sql.DB

// ConnectDB opens a connection to the database
func ConnectDB() {

	host := "localhost"
	port := 5432
	user := "postgres"
	password := "groot"
	dbname := "users"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	DB = db

}
