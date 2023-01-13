package main

import "fmt"

func main(){
	a:="Ashish"
	b := a[len(a)-1:]+a[1:len(a)-1]+a[0:1]

	fmt.Print(b)
}