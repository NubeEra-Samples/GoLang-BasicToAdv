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
	password = "Aditya@123"
	dbName   = "d1"
)

var db *sql.DB

// This function will make a connection to the database only once.
func main() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err = sql.Open("postgres", connStr)

	CheckError(err)

	err = db.Ping()

	CheckError(err)
	fmt.Println("The database is connected")
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
