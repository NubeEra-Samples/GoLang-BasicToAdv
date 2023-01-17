package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("storingData.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(fileName string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(fp *os.File) {
	fmt.Println("Writting")
	fmt.Fprintln(fp, "Welcome")
}
func closeFile(fp *os.File) {
	fmt.Println("Closing")
	err := fp.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v \n", err)
		os.Exit(1)
	}
}
