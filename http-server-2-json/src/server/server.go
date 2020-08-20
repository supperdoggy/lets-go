package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
)

var cache = make(map[string]worker)

const (
	htmlFormType = "application/x-www-form-urlencoded"
	jsonFormType = "application/json"
)

// returns random string of length
func randomStringGenerator(l int) string {
	// letters and number to random pick
	const lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	result := make([]byte, l)
	for i := range result {
		result[i] = lettersAndNumbers[rand.Intn(len(lettersAndNumbers))]
	}
	return string(result)
}

func parseForm(req *http.Request) (err error) {
	err = req.ParseForm()
	if err != nil {
		return
	}
	return
}

func getJsonData(req *http.Request) (url.Values, error) {
	// result
	mymap := make(map[string]string)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
		return url.Values{}, err
	}
	err = json.Unmarshal(body, &mymap)
	if err != nil {
		fmt.Println(err.Error())
		return url.Values{}, err
	}
	return jsonToUrlValues(&mymap), nil
}

// checking content type
// return ether url.values ether map
func handlePostRequest(req *http.Request) (url.Values, error) {
	switch req.Header.Get("content-type") {
	// html form
	case htmlFormType:
		parseForm(req)
		return req.Form, nil
		// json
	case jsonFormType:
		return getJsonData(req)
		// other
	default:
		return nil, fmt.Errorf("wrong request type")
	}
}

func jsonToUrlValues(mymap *map[string]string) url.Values {
	values := make(url.Values)
	for k, v := range *mymap {
		values.Add(k, v)
	}
	return values
}

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
		fmt.Println("get worker request")
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
