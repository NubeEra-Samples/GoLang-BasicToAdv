package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("storingData.txt")
	if err != nil {
		log.Fatalf("Unable to Open File %v", err)
	}
	buf := make([]byte, 1024)
	n, err := f.Read(buf)
}
