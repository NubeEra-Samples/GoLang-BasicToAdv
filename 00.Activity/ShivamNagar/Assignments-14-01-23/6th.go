package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	s := a * a * a
	fmt.Println(s)

}
