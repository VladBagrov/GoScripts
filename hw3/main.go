package main

import "fmt"

const (
	w = 8
	h = 8
)

func main() {
	var r rune
	for i := range w {
		for j := range h {
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
