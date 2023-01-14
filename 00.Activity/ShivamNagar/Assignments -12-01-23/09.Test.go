package main

import "fmt"

func main() {
	// Ques 1
	s := "Shivam"
	fmt.Println(s)
	var z = s[0]
	var x = s[len(s)-1]
	s = string(x) + s[1:len(s)] + string(z)
	fmt.Println(s)

	// Ques 2
	q := "Shivam"
	fmt.Println(q)
	var n int = 2
	q = q[:n] + q[n+1:]
	fmt.Println(q)
	// Ques 3
	var ll = " Hello World"
	fmt.Println(ll)
	for i := 0; i < len(ll); i++ {
		if ll[i] == ' ' {
			ll = ll[0:i] + "_" + ll[i+1:]
		}
	}
	fmt.Println(ll)

	// QUes 4
	check := "Hello World"
	countWords := 0
	countChar := 0
	for i := 0; i < len(check); i++ {
		if check[i] == ' ' {
			countWords = countWords + 1
		} else {
			countChar = countChar + 1
		}
	}
	fmt.Println(countWords + 1)
	fmt.Println(countChar)

}
