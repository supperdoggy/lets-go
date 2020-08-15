package main

import (
	"fmt"
	"sync"
)

func workerToHuman(w *worker) human {
	return w.human
}

func printWorkers(m map[int]worker, c chan bool, group *sync.WaitGroup) {
	defer group.Done()
	if <-c {
		fmt.Println("Starting printing workers")
		for _, v := range m {
			if v.isBoss != true {
				fmt.Println(v)
			}

		}
		fmt.Println("Done printing workers")
	}

}

func printBosses(m map[int]worker, c chan bool) {
	fmt.Println("Printing bosses")
	for _, v := range m {
		if v.isBoss {
			fmt.Println(v)
		}
	}
	fmt.Println("Done printing bosses")
	c <- true
}

func main() {
	// before transforming
	worker1 := worker{human: human{name: "David", age: "20"},
		job:    "Singer",
		isBoss: false}

	fmt.Printf("%T, %v\n", worker1, worker1)
	// after transforming
	human1 := workerToHuman(&worker1)
	fmt.Printf("%T, %v\n", human1, human1)

	w2 := worker{
		job:    "Cool boss",
		isBoss: true,
		human:  human{},
	}

	w3 := worker{
		job:    "Cool guy",
		isBoss: false,
		human:  human{},
	}

	w4 := worker{
		job:    "Programmer",
		isBoss: false,
		human:  human{},
	}

	m := make(map[int]worker)

	m[0] = worker1
	m[1] = w2
	m[2] = w3
	m[3] = w4

	c := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go printWorkers(m, c, &wg)
	go printBosses(m, c)

	wg.Wait()
}
