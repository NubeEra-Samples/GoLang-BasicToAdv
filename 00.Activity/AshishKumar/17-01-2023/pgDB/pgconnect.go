package main

import (
	"database/sql"
	"fmt"
    "github.com/spf13/viper"
	_ "github.com/lib/pq"
)

var db *sql.DB

// This function will make a connection to the database only once.
func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	var err error
	//connStr := "postgres://postgres:password@localhost/DB_1?sslmode=disable"
	connStr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", viper.Get("host"), viper.Get("port"), viper.Get("user"), viper.Get("password"), viper.Get("dbName"))
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
