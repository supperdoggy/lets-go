package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const(
	htmlFormType = "application/x-www-form-urlencoded"
	jsonFormType = "application/json"
)

// getting data from html form
func getFormData(req *http.Request) (url.Values, error) {
	if err := req.ParseForm(); err != nil {
		return url.Values{}, err
	}
	return req.Form, nil
}

// in case we got html form data
func handleFormData(req *http.Request) (answer string, err error) {
	// in case we get an error parsing form returning an error
	form, err := getFormData(req)
	if err != nil {
		fmt.Println(err.Error())
		answer = "Got an error parsing your form"
		return
	}

	switch form.Get("req") {
	case "ping":
		answer = "pong"
	case "pong":
		answer = "ping"
	default:
		answer = "wrong request"
	}
	return
}

// in case we ge
func handleJsonData(req *http.Request) (answer string, err error){
	body, err := ioutil.ReadAll(req.Body)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	mymap := make(map[string]string)
	err = json.Unmarshal(body, &mymap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch mymap["req"] {
	case "ping":
		answer = "pong"
	case "pong":
		answer = "ping"
	default:
		answer = "wrong request"
	}
	return
}

// checking content type
func handlePostRequest(req *http.Request) (answer string, err error){
	switch req.Header.Get("content-type") {
	// html form
	case htmlFormType:
		fmt.Println("got form request")
		return handleFormData(req)
		// json
	case jsonFormType:
		fmt.Println("got json request")
		return handleJsonData(req)
		// other
	default:
		return "wrong data format", nil
	}
}

// ping pong request handling
func pingPong(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		return
	case "POST":
		// getting answer
		answer, err := handlePostRequest(req)
		if err != nil {
			fmt.Println(err)
		}
		// sending answer
		_, _ = fmt.Fprint(writer, answer)
	}
}

func main() {
	http.HandleFunc("/", pingPong)
	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
