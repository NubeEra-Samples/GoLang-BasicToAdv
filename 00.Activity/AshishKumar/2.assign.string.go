package main

import "fmt"

func main(){
	a:="Hello there"
	n:=4
	a = a[:n]+a[n+1:]

	fmt.Print(a)
}