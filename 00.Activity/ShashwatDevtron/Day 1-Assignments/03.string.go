package main

import (
	"fmt"
)

func main() {
	s := "Hello I am Shashwat"
	n := []rune(s)
	// delete 5th character
	temp := append(n[:5], n[6:]...)
	fmt.Printf("%s",string(temp))
}