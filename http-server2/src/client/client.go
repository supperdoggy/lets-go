package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// creating new worker on server
func newWorker(name, position, job, email, phone, age string) {
	// url
	const u = "http://localhost:8080/newWorker"
	// creating worker
	formData := url.Values{
		"name":     {name},
		"position": {position},
		"job":      {job},
		"email":    {email},
		"phone":    {phone},
		"age":      {age},
	}
	// getting response
	_, err := http.PostForm(u, formData)
	// if we get error then print it and return nil
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// getting info about worker
func getWorkerAbout(id string) {
	// url
	const u = "http://localhost:8080/getWorker"
	// data for POST request
	formData := url.Values{
		"id": {id},
	}
	// getting response
	_, err := http.PostForm(u, formData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	newWorker("David", "Junior", "Programmer", "supperdoggy@knu.ua", "+380951102363", "18")
	getWorkerAbout("BpLnfgDsc3WD9F3qNfHK6a95")
}
