package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id   int    `json:"Id"`
	Name string `json:"NameOfProduct"`
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "welcome")
	list := Product{
		45,
		" Aditya",
	}
	json.NewEncoder(w).Encode(list)

}

func main() {

	http.HandleFunc("/", indexPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
