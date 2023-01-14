package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Scanln(&n)
	for i:=2;i<=n;i+=2{
		fmt.Printf("%v ",i)
	}
}