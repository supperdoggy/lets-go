package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func postRequest(u string, val url.Values) {
	resp, err := http.PostForm(u, val)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	postRequest("http://localhost:8080/", url.Values{"req": {"ping"}})
	postRequest("http://localhost:8080/", url.Values{"req": {"pong"}})
}
