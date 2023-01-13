package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Name string
	Id   int
}

var Products = []Product{
	{Name: "Mac", Id: 1},
	{Name: "Mac 2", Id: 2}}

func getProducts(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, Products)
	json.NewEncoder(w).Encode(Products)

}
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "%v", "Show List by Click this =<a href='/list'> Show List</a>")

	// fmt.Fprintf(w,)

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/list", getProducts)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
