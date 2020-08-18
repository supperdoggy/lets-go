package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func ping(w http.ResponseWriter, req *http.Request){
	switch req.Method {
	case "GET":
		return
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		if p := req.Form.Get("ping"); p=="ping"{
			fmt.Println("pong")
			postRequest("pong")

		}else{
			fmt.Println("Error")
		}

	}
}

func pong(w http.ResponseWriter, req *http.Request){
	switch req.Method {
	case "GET":
		return
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		if p := req.Form.Get("pong"); p=="pong"{
			fmt.Println("ping")
			postRequest("ping")
		}else{
			fmt.Println("Error")
		}

	}
}

func postRequest(s string){
	formData := url.Values{
		s:{s},
	}
	var url string = "http://localhost:8090/"+s
	_, err := http.PostForm(url, formData)
	if err != nil{
		fmt.Println(err)
	}
}

func start()  {
	postRequest("ping")
}

func main() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/pong", pong)
	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
