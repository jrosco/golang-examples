package main

import "fmt"

func main() {
	// An array
	a := [5]int{1, 2, 3, 4, 5}

	// A slice
	s := []string{"one", "two", "three"}

	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", s)

	newS := append(s, "four")

	e := newS[3]
	l := len(newS)

	fmt.Println(e, l)
}
