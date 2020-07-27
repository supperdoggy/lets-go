package main

import (
	"fmt"
)

// man
type man struct {
	name string
	age  uint
	sex  string
}

// worker
type worker struct {
	man
}

// jobs begin
type programmer struct {
	worker
	salary     uint
	position   string
	company    string
	freeCoffee bool
}

type designer struct {
	worker
	salary     uint
	position   string
	company    string
	freeCoffee bool
}

type sportsman struct {
	worker
	salary uint
	sport  string
	rank   string
	team   string
}

// jobs end

func main() {
	// sportsman
	first := sportsman{worker: worker{man: man{name: "David", age: 25, sex: "male"}},
		salary: 1300,
		sport:  "Basketball",
		rank:   "first",
		team:   "Lakers"}
	// designer
	second := designer{worker: worker{man: man{name: "Dasha", age: 40, sex: "female"}},
		salary:     2020,
		position:   "head designer",
		company:    "DesComp",
		freeCoffee: true}
	// programmer
	third := programmer{worker: worker{man: man{name: "Lesha", age: 22, sex: "male"}},
		salary:     34231,
		position:   "junior programmer",
		company:    "freelance",
		freeCoffee: true}

	fmt.Println(first)  // sportsman
	fmt.Println(second) // designer
	fmt.Println(third)  // programmer

	/*		OUTPUT
	{{{David 25 male}} 1300 Basketball first Lakers}
	{{{Dasha  40 female}} 2020 head designer DesComp true}
	{{{Lesha 22 male}} 34231 junior programmer freelance true}
	*/
}
