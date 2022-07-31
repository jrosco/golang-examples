package main

import "fmt"

type person struct {
	name   string
	age    int
	skills string
}

func main() {
	p := person{"Kasper", 33, "Mathematics and programming"}

	fmt.Println(p.age)
	fmt.Println(p.name)
	fmt.Println(p.skills)

}
