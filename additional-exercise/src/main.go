package main

import "fmt"

// cache for information about professions
var (
	infoAboutWorkers = make(map[WorkerPrinter]string)
)

// function gets arguments and outputs its types
func getType(a ...interface{}) {
	for _, v := range a {
		fmt.Printf("%T\n", v)
	}
}

func main() {
	// creating interfaces
	var w1 WorkerPrinter = &Programmer{Position: "TeamLead", Worker: Worker{Name: "Daniel Popov"}}
	var w2 WorkerPrinter = &BasketBallPlayer{Position: "PG", Worker: Worker{Name: "Max Maric"}}
	var w3 WorkerPrinter = &Singer{Position: "LeadSinger", Worker: Worker{Name: "Marlin Manson"}}

	// creating array of interfaces
	var a [3]WorkerPrinter
	a[0] = w1 // programmer
	a[1] = w2 // basketball player
	a[2] = w3 // singer

	// Writing results of getInfo() into cache
	for _, s := range a {
		infoAboutWorkers[s] = s.getInfo()
	}

	// outputting caches keys and values
	for k, v := range infoAboutWorkers {
		fmt.Println("Key:", k, ". Value:", v)
	}

	// outputting types of arguments
	getType(12, 223., "string", infoAboutWorkers)
}

// OUTPUT
//
// Key: &{TeamLead  {Daniel Popov}} . Value: Daniel Popov, Programmer : TeamLead
// Key: &{PG {Max Maric}} . Value: Max Maric, BasketBall player : PG
// Key: &{LeadSinger {Marlin Manson}} . Value: Marlin Manson, Singer : LeadSinger
// int
// float64
// string
// map[main.WorkerPrinter]string
//
