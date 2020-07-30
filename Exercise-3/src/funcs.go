package main

func changeName(p *human, name string) {
	*p = human{name: name, age: p.age}
}
