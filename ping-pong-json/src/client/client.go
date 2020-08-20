package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const(
	jsonContentType = "application/json"
	u = "http://localhost:8080/"
)

// reading plain text answer got from server
func readPlainTextResponse(response *http.Response) (string, error){
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(body), nil
}

// POST request method html form
func postRequestForm(u string, val url.Values) (answer string, err error){
	// sending POST request to server
	response, err := http.PostForm(u, val)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// getting answer from server
	answer, err = readPlainTextResponse(response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// returning answer
	return
}

// POST request method json
func postRequestJson(u string, val map[string]string) (answer string, err error){
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
	// getting answer from server
	answer, err = readPlainTextResponse(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// returning answer
	return
}

func main() {
	// HTML POST
	fmt.Println(postRequestForm(u, url.Values{"req": {"ping"}}))
	fmt.Println(postRequestForm(u, url.Values{"req": {"pong"}}))
	// JSON POST
	mymap := make(map[string]string)
	mymap["req"] = "ping"
	fmt.Println(postRequestJson(u, mymap))
	mymap["req"] = "pong"
	fmt.Println(postRequestJson(u, mymap))

}
