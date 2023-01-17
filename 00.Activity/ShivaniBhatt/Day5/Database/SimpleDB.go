package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("mysql", "root:123@tcp(192.168.1.88:3306)/a1")

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
