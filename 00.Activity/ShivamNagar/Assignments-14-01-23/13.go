package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	var lastDigit int
	for a > 0 {
		lastDigit = a % 10
		a = a / 10
	}
	fmt.Println(lastDigit)
}
