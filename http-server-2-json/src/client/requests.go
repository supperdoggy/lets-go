package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// making post request with html form values
// "application/x-www-form-urlencoded"
func htmlFormPostRequest(u string, values url.Values) (answer interface{}, err error) {
	// getting response
	resp, err := http.PostForm(u, values)
	// if we get error then print it and return nil
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return responseReader(resp)
}

// POST request method json
// "application/json"
func postRequestJson(u string, val map[string]string) (answer interface{}, err error) {
	// encoding values
	out, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(out)
	// sending POST request to server
	resp, err := http.Post(u, jsonContentType, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// returning answer
	return responseReader(resp)
}

func responseReader(response *http.Response) (interface{}, error) {
	switch response.Header.Get("Content-Type") {
	case jsonContentType:
		return readJsonResponse(response)
	default:
		return readPlainTextResponse(response)
	}
}
