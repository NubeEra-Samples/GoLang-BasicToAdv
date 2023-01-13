// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	r1 := 'B'
// 	r2 := 'b'
// 	r3 := '\a'

// 	fmt.Printf("R1: %c ", r1)
// 	fmt.Printf("R2: %U ", r2)
// 	fmt.Printf("R3: %s ", reflect.TypeOf(r3))
// }

package main

import (
	"fmt"
)

func main() {
	r1 := 'B'
	r2 := 'b'
	r3 := '\a'
	r4 := 4.5

	fmt.Printf("R1: %c ", r1)
	fmt.Printf("R2: %U ", r2)
	fmt.Printf("R3: %T ", r3)
	// fmt.Printf("R3: %s ", reflect.TypeOf(r3))
	// fmt.Println(reflect.TypeOf(r4))
}
