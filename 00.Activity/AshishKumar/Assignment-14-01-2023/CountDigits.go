package main

import (
	"fmt"
)

func main(){
	n:=123424
	count:=0
	for n>0{
		n=n/10
		count++
	}
	fmt.Printf("The number of digits is %v",count)
}