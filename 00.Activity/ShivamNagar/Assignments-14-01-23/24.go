package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	var product float64 = 1.0
	var ld float64
	for a > 0 {
		ld = (float64)(a % 10)
		product = product * ld
		a = a / 10
	}
	fmt.Println(product)

}
