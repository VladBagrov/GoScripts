package main

import (
	"fmt"
)

func main() {
	s := "Gopher"

	fmt.Printf("Hello, world %s\n", s)
	sum := aPlusB(5, 19)
	fmt.Print(sum)
}

func aPlusB(a, b int) int {
	return a + b
}
