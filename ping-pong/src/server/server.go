package main

import (
	"fmt"
	"log"
	"net/http"
)
// getting request
func sendAnswer(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		return
	case "POST":
		// parsing request form
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(writer, "ParseForm() err: %v", err)
			return
		}
		// checking request we got
		// if its ping sending pong
		// if its pong sending ping
		// else sending error
		var answer string
		switch req.Form.Get("req") {
		case "ping":
			answer = "pong"
		case "pong":
			answer = "ping"
		default:
			answer = "wrong request"
		}
		// printing answer we send to user
		fmt.Println(answer)
		// sending answer
		fmt.Fprint(writer, answer)
	}
}

func main() {
	http.HandleFunc("/", sendAnswer)
	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
