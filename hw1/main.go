package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text")
	s, err := reader.ReadString('\n')
	// s, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(s))
	fmt.Printf("dfsdfdsf %s", string(s))
	fmt.Printf("edewdwed")
}
