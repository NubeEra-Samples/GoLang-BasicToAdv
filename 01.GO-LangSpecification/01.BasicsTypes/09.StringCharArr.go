package main

import "fmt"

func main() {
	str := "abc"
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		c := string(chars[i])
		println(c)
	}

	//fmt.Println("Welcome"[1])
	//fmt.Printf("%c", "Welcome"[1])
	fmt.Printf("%c", []rune("Welcome")[1])
}
