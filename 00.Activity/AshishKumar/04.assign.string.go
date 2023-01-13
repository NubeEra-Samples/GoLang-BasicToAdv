package main

import (
	"fmt"
)

func main(){
    a:="My name is Ashish"
	countc:=0
	countw:=1
	for i:=0;i<len(a);i++ {
		if(a[i]==' '){
           countw++
		   continue
		}
		countc++
	}
	fmt.Print("No of words \n",countw)
	fmt.Print("No of characters ",countc)
}