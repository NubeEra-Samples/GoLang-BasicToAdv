package main

import (
	"fmt"
)

func main(){
	var n int
	fmt.Scanln(&n)
	if(n==0){
		fmt.Printf("1")
		return
	}
	fact:=1
	for i:=1;i<=n;i++{
		fact = fact*i
	}
	fmt.Printf("The factorial of %v is %v",n,fact)
}