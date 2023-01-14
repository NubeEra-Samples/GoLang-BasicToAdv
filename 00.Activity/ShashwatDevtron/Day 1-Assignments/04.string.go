package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello I am Shashwat"
	// delete 5th character
	temp := strings.Replace(s," ", "-", -1)
	fmt.Printf("%s",string(temp))
}