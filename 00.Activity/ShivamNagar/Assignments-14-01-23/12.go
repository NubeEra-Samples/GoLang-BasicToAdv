package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fact(n-1) * n
}
func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	fmt.Println(fact(a))

}
