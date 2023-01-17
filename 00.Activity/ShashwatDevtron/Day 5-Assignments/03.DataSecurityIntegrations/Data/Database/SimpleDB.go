package main

import (
	"database/postgres"
	"fmt"

	_ "github.com/go-sql-driver/postgres"
)

func main() {

	db, err := postgres.Open("postgres", "root:123@tcp(127.0.0.1:5432)/d1?ssmode=disable")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// INSERT INTO T1(Id,Name) VALUES(3,'CCC')
	fmt.Println("Inserting Data")
	insert, err := db.Query("INSERT INTO T1(Id,Name) VALUES(3,'CCC')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
