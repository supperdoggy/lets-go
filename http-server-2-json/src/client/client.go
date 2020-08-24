package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	newWorkerUrl    = "http://localhost:8080/newWorker"
	getWorkerUrl    = "http://localhost:8080/getWorker"
	jsonContentType = "application/json"
	jsonGetUrl      = "http://localhost:8080/getJson"
	htmlFormType    = "application/x-www-form-urlencoded"
	plainTextType   = "text/plain"
)

// reading plain text answer got from server
func readPlainTextResponse(response *http.Response) (string, error) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(body), nil
}

func readJsonResponse(response *http.Response) (answer map[string]string, err error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = json.Unmarshal(body, &answer)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func main() {
	// HTML FORM
	//
	// new worker
	formData := url.Values{
		"name":     {"David"},
		"position": {"Junior"},
		"job":      {"Programmer"},
		"email":    {"one@one.go"},
		"phone":    {"+380951102363"},
		"age":      {"18"},
	}
	// making html POST request and getting answer
	answer, err := htmlFormPostRequest(newWorkerUrl, formData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(answer)
	// get worker
	formData = url.Values{
		"id": {"BpLnfgDsc3WD9F3qNfHK6a95"},
	}
	answer, err = htmlFormPostRequest(getWorkerUrl, formData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// printing answer
	fmt.Println(answer)
	// JSON
	//
	// new worker
	mymap := map[string]string{
		"name":     "Max",
		"position": "Head",
		"job":      "Designer",
		"email":    "lol@knu.ua",
		"phone":    "+380951112363",
		"age":      "22",
	}
	// making json POST request and getting plain text answer
	answer, err = postRequestJson(newWorkerUrl, mymap)
	// printing answer
	fmt.Println(answer)
	// getWorker
	mymap = map[string]string{
		"id": "jjJkwzDkh0h3fhfUVuS0jZ9u",
	}
	// same, making request
	answer, err = postRequestJson(getWorkerUrl, mymap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// printing answer we got
	fmt.Println(answer)

	// get json answer
	resp, err := http.Get(jsonGetUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	mymap, err = readJsonResponse(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(mymap)

}
