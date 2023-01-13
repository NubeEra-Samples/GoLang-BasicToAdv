package main

import (
	"fmt"
	"log"
	"net/http"
)

type  Product struct{
	Id  string
	Name string
}


func indexPage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
   fmt.Fprint(w,"Welcome")

   fmt.Fprintf(w,"%v",`<a href='/list' >Show List</a>`)
}

func showProducts(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/html")

	data := []Product{{"Id", "Name"},{"1","abc"},{"2","def"}}
    code:="";
	for _,val:= range data {
		code+="<tr style='border:1px solid black;'><td style='border:1px solid black;'>"+val.Id+"</td><td style='border:1px solid black;'>"+val.Name+"</td><tr>"
	}
	code= "<table style='border:1px solid black;width:'50%'' >"+code+"</table>"
	fmt.Fprintf(w,"%v",code)
}

func main(){
     http.HandleFunc("/index",indexPage)
     http.HandleFunc("/list",showProducts)
	 log.Fatal(http.ListenAndServe(":8080",nil))
}