package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id   int    `json:"id"`
	Name string `josn:"name"`
}

var products []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome to the Web Page\n")
	fmt.Fprintf(w, "LIST\n")
	json.NewEncoder(w).Encode(products)
}
func getFrontPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Web Page")

}

func main() {
	products = append(products, Product{Id: 1, Name: "Product1"})
	products = append(products, Product{Id: 2, Name: "Product2"})

	http.HandleFunc("/getProducts", getProducts)
	http.HandleFunc("/getFirstPage", getFrontPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
