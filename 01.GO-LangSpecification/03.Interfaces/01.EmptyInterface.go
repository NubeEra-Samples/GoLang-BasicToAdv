package main

import "fmt"

func observe(i interface{}) {
	fmt.Printf("The type passed is: %T\n", i)
	fmt.Printf("The value passed is: %#v \n", i)
	fmt.Println("-------------------------------------")
}
func main() {
	var value float64 = 15
	value2 := "testdatatest"
	observe(value)
	observe(value2)
}
