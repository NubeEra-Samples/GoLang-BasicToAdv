package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("Enter the Number:")
	fmt.Scanln(&a)
	if a%5==0 && a%11==0{
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}

}
