package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	var b int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&b)
	if a > b {
		fmt.Println(a)
	} else if b > a {
		fmt.Println(b)
	} else {
		fmt.Println("Equal")
	}

}
