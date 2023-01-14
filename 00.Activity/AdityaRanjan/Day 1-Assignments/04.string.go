package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "My name is Aditya"
	// delete 5th character
	temp := strings.Replace(s, " ", "-", -1)
	fmt.Printf("%s", string(temp))
}
