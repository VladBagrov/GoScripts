package main

import "fmt"

func main() {
	var w, h int
	fmt.Println("Enter width desk:")
	fmt.Scanf("%d\n", &w)

	fmt.Println("Enter height desk:")
	fmt.Scanf("%d\n", &h)

	var r rune
	for i := range h {
		for j := range w {
			r = what_r(i, j)
			fmt.Print(string(r))
		}
		fmt.Print("\n")
	}
}

func what_r(i, j int) rune {
	var s rune
	if (i%2 == 0 && j%2 == 0) || (i%2 == 1 && j%2 == 1) {
		s = '#'
	}
	if (i%2 == 0 && j%2 == 1) || (i%2 == 1 && j%2 == 0) {
		s = ' '
	}

	return s
}
