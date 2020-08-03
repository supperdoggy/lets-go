package main

import (
	"fmt"
)

type Programmer struct {
	Position string
	Job string
	Worker
}

func (p *Programmer) getInfo() {
	fmt.Println(p.Name, ",", "Programmer", p.Position)
}

type BasketBallPlayer struct {
	Position string
	Worker
}

func (b *BasketBallPlayer) getInfo() {
	fmt.Println(b.Name, ",", "BasketBall player :", b.Position)
}

type Singer struct {
	Position string
	Worker
}

func (s *Singer) getInfo() {
	fmt.Println(s.Name, ",", "Singer : ", s.Position)
}

type WorkerPrinter interface {
	getInfo()
}
