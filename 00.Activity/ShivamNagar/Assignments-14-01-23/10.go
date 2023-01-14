package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&n)
	for i := 2; i < n; i = i + 2 {
		fmt.Println(i)
	}

}
