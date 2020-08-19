package main

import (
	"fmt"
	"log"
	"net/http"
)

func getAnswer(writer http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		return
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(writer, "ParseForm() err: %v", err)
			return
		}

		var answer string
		switch req.Form.Get("req") {
		case "ping":
			answer = "pong"
		case "pong":
			answer = "ping"
		default:
			answer = "wrong request"
		}
		fmt.Println(answer)
		fmt.Fprint(writer, answer)
	}
}

func main() {
	http.HandleFunc("/", getAnswer)
	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
