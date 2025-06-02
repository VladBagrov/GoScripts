package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//var width int
	//var s string
	//
	//fmt.Print("Enter lenght\n")
	//_, err := fmt.Fscan(os.Stdin, &s)
	//if err != nil {
	//	log.Println("error reading input")
	//}
	//
	//fmt.Print("Enter width\n")
	//_, err = fmt.Fscan(os.Stdin, &width)
	//if err != nil {
	//	log.Println("error reading input")
	//}
	//
	//fmt.Printf("len = %s, width = %d\n", s, width)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text")
	// s, err := reader.ReadString('\n')
	s, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(s))
	fmt.Printf("dfsdfdsf %s", string(s))
	fmt.Printf("edewdwed")
	fmt.Printf("edewdwed")
	fmt.Printf("edewdwed")

}
