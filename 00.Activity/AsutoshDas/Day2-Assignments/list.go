package main

import (
	"fmt"
	"log"
	"net/http"
)

type product struct {
	Id string
	Name string
}
func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div><a href='%s'>%s</a></div>", "Welcome to Web with GO", "/list", "List Of Products")
}

func productList(w http.ResponseWriter, r *http.Request) {
	var products = []product {
		 {
			Id:"1",
			Name: "Product1",
		},
		 {
			Id:"2",
			Name: "Product2",
		},
		 {
			Id:"3",
			Name: "Product3",
		},
		 {
			Id:"4",
			Name: "Product4",
		},
	}
	responseHTML := `<table>
	<tr>
	  <th>Id</th>
	  <th>Name</th>
	</tr>`
  for _, p := range products{
	responseHTML += `<tr>
	<td>`+p.Id+`</td>
	<td>`+p.Name+`</td>
  </tr>`
  }
  responseHTML += `</table>`
	fmt.Fprintf(w, responseHTML)
}

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/list", productList)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
