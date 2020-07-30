package main

import "fmt"

type human struct {
	name string
	age  int
}

func (p *human) happyBirthday() {
	fmt.Println("Happy birthday " + p.name)
	p.age++
	fmt.Printf("You are %v\n", p.age)
}

func (p *human) changeName(name string) {
	fmt.Printf("Name changed to %v\n", name)
	p.name = name
}
