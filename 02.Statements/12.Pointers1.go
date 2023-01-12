package main

func zeroVal(iVal int) {
	iVal = 0
}
func zeroPtr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	println("Initial: ", i)

	zeroVal(i)
	println("After zeroVal Called: ", i)

	zeroPtr(&i)
	println("After zeroPtr Called: ", i)

	println("Pointer can call address of i : ", &i) // 0xHexaDecimalValue
}
