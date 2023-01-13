package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("Length: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Dcl: ", b)
	//2X3 --> Row X Column
	var twoD [2][3]int
	for r := 0; r < 2; r++ {
		for c := 0; c < 3; c++ {
			twoD[r][c] = r + c
		}
	}
	//[ [0 1 2] [1 2 3] ]
	fmt.Println("2d: ", twoD)
}

// [ [1,2,3] [3,2,1] ]
