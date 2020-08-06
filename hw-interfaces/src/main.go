package main

import "fmt"

var (
	// hash for information about professions
	infoAboutWorkers = make(map[WorkerPrinter]string)
	// hash for storing different values of different types
	differentTypes = make(map[interface{}]interface{})
)

// function gets map and outputs its types
func getType(a map[interface{}]interface{}) {
	for k, v := range a {
		fmt.Printf("Key: %v, value type :%T\n", k, v)
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

	// assigning different keys to different values
	differentTypes[0] = uint(2)
	differentTypes[1.] = false
	// int(2) != byte(2)
	differentTypes[2] = "Dave"
	differentTypes[byte(2)] = infoAboutWorkers

	differentTypes[4i+2] = differentTypes
	differentTypes["Hello world"] = "H"
	differentTypes[[3]int{}] = [3]int{}
	// outputting types of map
	getType(differentTypes)

}

// OUTPUT
//
//Key: &{LeadSinger {Marlin Manson}} . Value: Marlin Manson, Singer : LeadSinger
//Key: &{TeamLead  {Daniel Popov}} . Value: Daniel Popov, Programmer : TeamLead
//Key: &{PG {Max Maric}} . Value: Max Maric, BasketBall player : PG
//Key: (2+4i), value type :map[interface {}]interface {}
//Key: Hello world, value type :string
//Key: [0 0 0], value type :[3]int
//Key: 0, value type :uint
//Key: 1, value type :bool
//Key: 2, value type :string
//Key: 2, value type :map[main.WorkerPrinter]string
//
