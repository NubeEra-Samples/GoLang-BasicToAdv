package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	for i := a; i >= 1; i-- {
		fmt.Println(i)
	}

}
