package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.Form.Get("name")
		surname := r.Form.Get("surname")
		age := r.Form.Get("age")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Surname = %s\n", surname)
		fmt.Fprintf(w, "Age = %s\n", age)
	}
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Printf("Starting server...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
