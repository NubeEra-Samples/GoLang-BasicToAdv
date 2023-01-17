package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// This function will make a connection to the database only once.
func main() {
	var err error
	//connStr := "postgres://postgres:password@localhost/DB_1?sslmode=disable"
	connStr := "postgres://root:123@127.0.0.1/d1?sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
}
