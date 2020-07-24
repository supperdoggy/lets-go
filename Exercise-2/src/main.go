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

type streamer struct {
	worker
	salary      uint
	platform    string
	game        string
	subscribers uint
}

type musician struct {
	worker
	salary     uint
	band       string
	instrument string
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
	// musician
	second := musician{worker: worker{man: man{name: "Natalia", age: 30, sex: "female"}},
		salary:     5020,
		band:       "Cool band",
		instrument: "Guitar"}
	// streamer
	third := streamer{worker: worker{man: man{name: "Maks", age: 21, sex: "male"}},
		salary:      2000,
		platform:    "Twitch",
		game:        "dwarf fortress",
		subscribers: 1000}
	// designer
	fourth := designer{worker: worker{man: man{name: "Dasha", age: 40, sex: "female"}},
		salary:     2020,
		position:   "head designer",
		company:    "DesComp",
		freeCoffee: true}
	// programmer
	fifth := programmer{worker: worker{man: man{name: "Lesha", age: 22, sex: "male"}},
		salary:     34231,
		position:   "junior programmer",
		company:    "freelance",
		freeCoffee: true}

	fmt.Println(first)  // sportsman
	fmt.Println(second) // musician
	fmt.Println(third)  // streamer
	fmt.Println(fourth) // designer
	fmt.Println(fifth)  // programmer

	/*		OUTPUT
	{{{David 25 male}} 1300 Basketball first Lakers}
	{{{Natalia 30 female}} 5020 Cool band Guitar}
	{{{Maks 21 male}} 2000 Twitch dwarf fortress 1000}
	{{{Dasha  40 female}} 2020 head designer DesComp true}
	{{{Lesha 22 male}} 34231 junior programmer freelance true}
	*/
}
