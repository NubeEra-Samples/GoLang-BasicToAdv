package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	v1 := m["k1"]
	fmt.Println("Value in Map By Key: ", v1)
	fmt.Println("Lenght of MAP: ", len(m))

	m["k2"] = 25
	fmt.Println("Before Delete MAP: ", m)
	delete(m, "k2")
	fmt.Println("After Delete MAP: ", m)

	_, prs1 := m["k2"]
	fmt.Println("Catching k2 in MAP: ", prs1)

	_, prs2 := m["k1"]
	fmt.Println("Catching k1 in MAP: ", prs2)

}
