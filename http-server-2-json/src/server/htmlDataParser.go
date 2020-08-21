package main

import "net/http"

func parseForm(req *http.Request) (err error) {
	err = req.ParseForm()
	if err != nil {
		return
	}
	return
}
