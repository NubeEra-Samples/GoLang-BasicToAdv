package main

import (
	"fmt"
	"math"
)

func main(){
	var n int
	fmt.Scanln(&n)
    for i:=1;i<=int(math.Sqrt(float64(n)));i++{
         if(n%i==0){
			fmt.Printf("%v ",i)
		 }
	}
}