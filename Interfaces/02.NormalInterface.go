// Go program to illustrate the concept
// of dynamic values and types
package main

import "fmt"

// Creating an interface
type tank interface {

	// Methods
	Tarea() float64
	Volume() float64
}

func main() {

	var t tank
	fmt.Println("Value of the tank interface is: ", t)
	fmt.Printf("Type of the tank interface is: %T ", t)
}
