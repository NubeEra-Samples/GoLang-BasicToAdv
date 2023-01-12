package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("Sum of Array: ", sum)

	for i, n := range nums {
		if n == 3 {
			fmt.Println("Index: ", i)
		}
	}
}
