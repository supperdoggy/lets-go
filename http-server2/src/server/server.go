package main

import(
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}

