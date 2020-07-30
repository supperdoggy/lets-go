package main

import "fmt"

type man struct {
	name *string
	age  *int
}

func (m *man) happyBirthday() {
	var age int = *m.age + 1
	m = &man{name: m.name, age: &age}

	fmt.Println("Happy birthday " + *m.name)
	fmt.Printf("You are %v\n", *m.age)
}

func (m *man) changeName(name string) {
	fmt.Printf("Name changed to %v\n", name)
	m.name = &name
}
