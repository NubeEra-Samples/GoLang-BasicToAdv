package main

import (
	"fmt"
)

func main() {
	s := "abcdefgh"
	temp := []rune(s)
	// fmt.Println(temp)
	temp[0], temp[len(temp)-1] = temp[len(temp)-1] , temp[0]
	fmt.Printf("%s",string(temp))

}