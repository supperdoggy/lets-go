package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var cache = make(map[string]worker)

const (
	htmlFormType = "application/x-www-form-urlencoded"
	jsonFormType = "application/json"
)

// Creating new worker using data in Post form
func newWorker(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		fmt.Println("Got new worker request")
		resp, err := handlePostRequest(req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// creating new worker with data in form
		w := worker{name: resp.Get("name"),
			age:         resp.Get("age"),
			position:    resp.Get("position"),
			job:         resp.Get("job"),
			phoneNumber: resp.Get("phone"),
			email:       resp.Get("email"),
			id:          randomStringGenerator(24),
		}
		// putting worker into cache
		cache[w.id] = w
		fmt.Fprint(writer, "Worker created successfully!")
	default:
		fmt.Fprint(writer, "Not POST request")
		// in case request method is not POST just returns nil
		return
	}
}

// Shows data for worker using private id of worker
func getWorker(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		fmt.Println("Got worker request")
		resp, err := handlePostRequest(req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		id := resp.Get("id")
		// checking if id is valid, if not just returning nil and print error
		if id == "" {
			fmt.Fprint(writer, "id is not valid")
			return
		}
		// getting worker and if ok printing his info
		if w, ok := cache[id]; ok {
			fmt.Fprint(writer, w.getInfo())
		} else {
			fmt.Fprint(writer, "Nothing found by your id")
		}
		return
	default:
		return
	}
}

// to send json!!!
func sendJson(writer http.ResponseWriter, request *http.Request) {
	userJson, err := json.Marshal(map[string]string{"ok": "ok"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(userJson)
}

func main() {
	fmt.Println("Starting server...")
	// path for creating new worker
	http.HandleFunc("/newWorker", newWorker)
	// path for getting info about a worker by id of the worker
	http.HandleFunc("/getWorker", getWorker)
	// sending json
	http.HandleFunc("/getJson", sendJson)
	// staring server
	err := http.ListenAndServe(":8080", nil)
	// if we get an error it will print
	if err != nil {
		fmt.Println(err.Error())
	}
}
