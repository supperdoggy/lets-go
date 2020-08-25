package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const jsonContentType = "application/json"

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

func postRequestJson(u string, val map[string]string) (answer string, err error) {
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
	return readPlainTextResponse(resp)
}


func main(){
	mymap := make(map[string]string)
	mymap["username"] = "admin"
	mymap["pass"] = "admin"
	answer, err := postRequestJson("http://localhost:8080/getToken", mymap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(answer)

	mymap = map[string]string{"token":string(answer)}
	answer, err = postRequestJson("http://localhost:8080/ping", mymap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(answer)
}
