package main

import (
	"fmt"
	"math"
)

func main() {
	var p float64 = 10000
	var r float64 = 2
	var t float64 = 4
	var n float64 = 1
	var result float64 = p*(math.Pow(1+(r/t), n*t)) - p
	fmt.Println(result)
}
