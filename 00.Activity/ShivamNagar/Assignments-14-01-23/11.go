package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	for i := 2; i*i < a; i++ {
		if a%i == 0 {
			fmt.Println(i)
		}
	}

}
