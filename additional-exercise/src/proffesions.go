package main

type Worker struct {
	Name string
}

// programmer profession
type Programmer struct {
	Position string
	Job      string
	Worker
}

func (p *Programmer) getInfo() string {
	return p.Name + ", " + "Programmer : " + p.Position
}

// basketball player profession
type BasketBallPlayer struct {
	Position string
	Worker
}

func (b *BasketBallPlayer) getInfo() string {
	return b.Name + ", " + "BasketBall player : " + b.Position
}

// singer profession
type Singer struct {
	Position string
	Worker
}

func (s *Singer) getInfo() string {
	return s.Name + ", " + "Singer : " + s.Position
}

// worker interface
type WorkerPrinter interface {
	// prints name, profession and position
	getInfo() string
}
