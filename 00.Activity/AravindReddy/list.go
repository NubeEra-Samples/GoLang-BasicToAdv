package main

import (
	"fmt"
	"log"
	"net/http"
)

type data struct {
	id string
	name string
}
func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div><a href='%s'>%s</a></div>", "Welcome to Web with GO", "/list", "products")
}
var samples []data
func listofproducts(w http.ResponseWriter, r *http.Request) {
	sample:=data{"1","aravind"}
	sample1:=data{"2","raj"}
	samples=append(samples,sample)
	samples=append(samples,sample1)

	
	html := `<table>
	<tr>
	  <th>id</th>
	  <th>name</th>
	</tr>`
	for _, samp := range samples{
		html += `<tr>
		<td>`+samp.id+`</td>
		<td>`+samp.name+`</td>
	  </tr>`}
      html += `</table>`
	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/list", listofproducts)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

