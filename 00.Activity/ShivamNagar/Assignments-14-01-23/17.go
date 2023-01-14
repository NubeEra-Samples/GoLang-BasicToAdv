package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	if a%4 == 0 {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("No Leap Year")
	}

}
