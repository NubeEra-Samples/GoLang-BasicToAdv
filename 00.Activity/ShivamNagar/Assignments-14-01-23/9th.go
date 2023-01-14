package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	if a%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
