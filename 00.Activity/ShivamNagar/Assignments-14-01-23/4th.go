package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	s := a
	var count int
	for ; a > 0; a = a / 10 {
		count++
	}
	fmt.Printf("No of Digits in %v is %v\n", s, count)
}
