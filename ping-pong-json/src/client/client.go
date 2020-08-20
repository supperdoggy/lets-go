package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func readPlainTextResponse(response *http.Response) (string, error){
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(body), nil
}

func postRequestForm(u string, val url.Values) (answer string, err error){
	response, err := http.PostForm(u, val)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	answer, err = readPlainTextResponse(response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func postRequestJson(u string, val map[string]string) (answer string, err error){
	out, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(out)
	resp, err := http.Post(u, "application/json", reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	answer, err = readPlainTextResponse(resp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

const(
	u = "http://localhost:8080/"
)

func main() {
	fmt.Println(postRequestForm(u, url.Values{"req": {"ping"}}))
	fmt.Println(postRequestForm(u, url.Values{"req": {"pong"}}))
	mymap := make(map[string]string)
	mymap["req"] = "ping"
	fmt.Println(postRequestJson(u, mymap))
	mymap["req"] = "pong"
	fmt.Println(postRequestJson(u, mymap))

}
