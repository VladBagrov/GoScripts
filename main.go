package main

import (
	"fmt"
)

type User struct {
	name  string
	phone string
	id    int32
}

func main() {
	//s := "gopher"
	//u := User{name: "vlad", phone: "334455", id: 1}
	//su := []User{}

	//su = append(append(su, User{name: "dd"}), User{name: "ff"})

	a := "anysome"
	m := []int32{1, 2, 3}
	m[0] = 100

	chU(&a)
	f := []int32{}

	//s := User{phone: "wwww"}
	fmt.Printf("Hello and welcome, %v!\n", f)
	//for i := 1; i <= 5; i++ {
	//	fmt.Println("i =", 100/i)
	//}
}

func chU(u *string) {
	*u = "serg"
}
