package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// making post request to given url with given values
func postRequest(u string, val url.Values) {
	// sending post request + getting response
	resp, err := http.PostForm(u, val)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	// parsing body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}

func main() {
	postRequest("http://localhost:8080/ping", url.Values{"req": {"ping"}})
	postRequest("http://localhost:8080/ping", url.Values{"req": {"pong"}})
}
