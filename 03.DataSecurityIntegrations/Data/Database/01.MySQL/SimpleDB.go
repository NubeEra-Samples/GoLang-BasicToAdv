package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/a1")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// INSERT INTO T1(Id,Name) VALUES(3,'CCC')
	fmt.Println("Inserting Data")
	insert, err := db.Query("INSERT INTO t1(Id,Name) VALUES(3,'CCC')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}
