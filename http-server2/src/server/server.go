package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var cache = make(map[string]worker)

// returns random string of lenght l
func randomStringGenerator(l int) string {
	// letters and number to random pick
	const lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	result := make([]byte, l)
	for i := range result {
		result[i] = lettersAndNumbers[rand.Intn(len(lettersAndNumbers))]
	}
	return string(result)
}

// Creating new worker using data in Post form
func newWorker(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		fmt.Println("Got POST request")
		// parsing form and checking for errors
		if err := req.ParseForm(); err != nil{
			// if there is error - printing it
			fmt.Println(err.Error())
			return
		}
		// creating new worker with data in form
		w := worker{name: req.Form.Get("name"),
			age:         req.Form.Get("age"),
			position:    req.Form.Get("position"),
			job:         req.Form.Get("job"),
			phoneNumber: req.Form.Get("phone"),
			email:       req.Form.Get("email"),
			id:          randomStringGenerator(24),
		}

		// putting worker into cache
		cache[w.id] = w
		fmt.Println("Worker created successfully!")
		fmt.Println(w.getInfo(), "\n")
	default:
		fmt.Println("Not POST request")
		// in case request method is not POST just returns nil
		return
	}
}

// Shows data for worker using private id of worker
func getWorker(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		// parsing form and checking for errors
		if err := req.ParseForm(); err != nil {
			// if we get an error then printing it
			fmt.Println(err.Error())
		}
		// getting id from form
		id := req.Form.Get("id")
		// checking if id is valid, if not just returning nil and print error
		if id == "" {
			fmt.Println("id is not valid")
			return
		}
		// getting worker and if ok printing his info
		if w, ok := cache[id]; ok {
			fmt.Println("Found worker!")
			fmt.Println(w.getInfo())
			fmt.Fprint(writer, w.getInfo())
		} else {
			fmt.Fprint(writer, "Nothing found by your id")
		}
		return
	default:
		return
	}
}

func main() {
	fmt.Println("Starting server...")
	// path for creating new worker
	http.HandleFunc("/newWorker", newWorker)
	// path for getting info about a worker by id of the worker
	http.HandleFunc("/getWorker", getWorker)
	// staring server
	err := http.ListenAndServe(":8080", nil)
	// if we get an error it will print
	if err != nil {
		fmt.Println(err.Error())
	}
}
