package main

import (
	"fmt"
	"math"
)

func main(){
	p,r,n,t:=1000.0,10.0,5.0,10.0
	ci:=p*math.Pow((1+(r/n)),n*t)-p

	fmt.Printf("The compound intrest is %v",ci)
}