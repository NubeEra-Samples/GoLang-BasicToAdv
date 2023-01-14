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
	var c int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&c)

	if a > b {
		if a > c {
			fmt.Println(a)
		} else if c > a {
			fmt.Println(c)
		}

	} else if b > a {
		if a > c {
			fmt.Println(a)
		} else if c > a {
			fmt.Println(c)
		}
	}
}
