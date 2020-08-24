package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// checking content type
// return ether url.values ether map
func handlePostRequest(req *http.Request) (url.Values, error) {
	switch req.Header.Get("content-type") {
	// html form
	case htmlFormType:
		err := req.ParseForm()
		if err != nil {
			return nil, err
		}
		return req.Form, nil
		// json
	case jsonFormType:
		return getJsonData(req)
		// other
	default:
		return nil, fmt.Errorf("wrong request type")
	}
}
