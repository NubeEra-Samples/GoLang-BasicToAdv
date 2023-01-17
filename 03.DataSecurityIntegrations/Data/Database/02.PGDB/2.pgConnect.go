package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "shivam@devtron.ai"
	dbName   = "dbshivam"
)

var db *sql.DB

// This function will make a connection to the database only once.
func main() {
	var err error
	//connStr := "postgres://postgres:password@localhost/DB_1?sslmode=disable"
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err = sql.Open("postgres", connStr)

	CheckError(err)

	err = db.Ping()

	CheckError(err)

	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
