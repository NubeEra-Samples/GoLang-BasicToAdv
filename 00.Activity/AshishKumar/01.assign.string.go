package main

import (
	"fmt"
)

func main(){
    a:="My name is Ashish"
	for i:=0;i<len(a);i++ {
		if(a[i]==' '){
            a= a[:i] + "-" +a[i+1:];
		}
	}
	fmt.Print(a)
}