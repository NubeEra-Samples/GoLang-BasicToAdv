package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	
	var s = add(2, 3)
	fmt.Printf("Result is %v", s)
}
