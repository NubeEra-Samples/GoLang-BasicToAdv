package main

func zeroVa1(iVa1 int) {
	iVa1 = 0
}
func zeroPtr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	println("lnitial: ", i)
	zeroVa1(i)
	println("After zeroVa1 Called: ", i)
	zeroPtr(&i)
	println("After zeroPtr Called: ", i)

}