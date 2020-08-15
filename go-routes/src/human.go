package main

type human struct {
	name string
	age  string
}

func (h *human) String() string {
	return h.name + ", " + string(h.age)
}
