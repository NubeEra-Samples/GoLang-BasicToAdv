package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello I am Shashwat"
	n := strings.Split(s," ")
	fmt.Println(len(n))	
	fmt.Println(len(s)-len(n)+1)	

}