package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// func send(msg string) {
// 	fmt.Println(msg)
// }

func main() {

	// nextInt := intSeq()

	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// newInts := intSeq()
	// fmt.Println(newInts())

	//send("Welcome to GO Lang")

	func(msg string) {
		fmt.Println(msg)
	}("Welcome to GO Lang")
}
