package main

func main() {
	var w1 WorkerPrinter = &Programmer{Position: "TeamLead", Worker: Worker{Name: "Daniel Popov"}}
	var w2 WorkerPrinter = &BasketBallPlayer{Position: "PG", Worker: Worker{Name: "Max Maric"}}
	var w3 WorkerPrinter = &Singer{Position: "LeadSinger", Worker: Worker{Name: "Marlin Manson"}}

	var a [3]WorkerPrinter
	a[0] = w1
	a[1] = w2
	a[2] = w3

	for _, s := range a {
		s.getInfo()
	}

}
