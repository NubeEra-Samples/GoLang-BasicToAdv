package main

import "fmt"

func send() func(string) {
	i := 1
	return func(msg string) {
		fmt.Println(msg)
	}
	i++
}

func main() {

	sf := send()
	sf("Welcome to GO Language")
}

//   obj1.f1().obj2.f2().{...}
