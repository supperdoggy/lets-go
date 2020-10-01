package main

type apiResponse struct {
	Ok bool `json:"ok"`
	Error string `json:"error"`
	Answer interface{} `json:"answer"`
}
