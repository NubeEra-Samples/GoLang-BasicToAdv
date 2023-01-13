package main

import (
	"fmt"
	"log"
	"net/http"
	// "html/template"
)

type product struct {
	Id string
	SuperHeros string
}

// var tmpl *template.template

var products = []product {
	{
	 Id:"1",
	 SuperHeros: "Terminator",
 },
	{
	 Id:"2",
	 SuperHeros: "SuperMan",
 },
	{
	 Id:"3",
	 SuperHeros: "Gladiator",
 },
}


func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div><a href='%s'>%s</a></div>", "Welcome to Web with GO", "/list", "List Of Products")

}

func superHeroList(w http.ResponseWriter, r *http.Request) {

	renderSupermanListResponse := `<table >
	<tr>
	  <th>Id</th>
	  <th>SuperHeros</th>
	</tr>`

  for _, p := range products{
	renderSupermanListResponse += `<tr>
	<td>`+p.Id+`</td>
	<td>`+p.SuperHeros+`</td>
  </tr>`
  }
  renderSupermanListResponse += `</table>`
	renderButton := fmt.Fprintf(w, "<div><a href='%s'>%s</a></div>", "/", "Go Back")
	renderSupermanListResponse += renderButton

	fmt.Fprintf(w, renderSupermanListResponse)
}

func main() {
	// tmpl = template.Must(template.ParseFile("index.gohtml"))
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/list", superHeroList)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
