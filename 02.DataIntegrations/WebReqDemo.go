package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Web with GO")
	fmt.Println("Data Transfering to Console ")
}

func main() {
	http.HandleFunc("/", indexPage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

