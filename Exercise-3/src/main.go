package main

import (
	"fmt"
)

func main() {
	p := human{name: "Dave", age: 14}
	fmt.Println(p.age)

	p.happyBirthday()
	fmt.Println(p.age)

	p.changeName("David")
	p.happyBirthday()

	changeName(&p, "Dave")
	p.happyBirthday()

	h := &p
	h.happyBirthday()

	age := 15
	name := "Max"
	m := man{age: &age, name: &name}
	m.happyBirthday()

	m.changeName("Maks")
	m.happyBirthday()
}
