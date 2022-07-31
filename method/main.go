package main

import "fmt"

type person struct {
	name   string
	age    int
	skills string
}

func (p *person) setSkill(skill string) {
	p.skills = skill
}

func main() {
	p := person{"Kasper", 33, "Mathematics and programming"}

	fmt.Println(p.age)
	fmt.Println(p.name)
	fmt.Println(p.skills)

	p.setSkill("Coding Go")

	fmt.Println("New skill:", p.skills)

}
