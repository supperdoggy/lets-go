package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
const (
	htmlFormType = "application/x-www-form-urlencoded"
	jsonFormType = "application/json"
)
var users = make(map[string]string)

func getSha256(text string) string {
	hashser := sha256.New()
	hashser.Write([]byte(text))
	return hex.EncodeToString(hashser.Sum(nil))
}

func register(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		fmt.Println("Got get request")
		http.ServeFile(writer, request, "reg.html")
	case "POST":
		fmt.Println("Got post request")
		if err := request.ParseForm(); err != nil {
			fmt.Println(err.Error())
			return
		}
		username := request.Form.Get("login")
		pass := getSha256(request.Form.Get("pass"))
		users[username] = pass
		fmt.Fprint(writer, "Registered")
	}
}

func validateUser(username, password string) bool{
	if pass, ok := users[username]; ok {
		if pass == getSha256(password) {
			return true
		}
	}
	return false
}


func login(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		fmt.Println("Got get request")
		http.ServeFile(writer, request, "login.html")
	case "POST":
		var answer string

		switch request.Header.Get("content-type") {
		case htmlFormType:
			if err := request.ParseForm(); err != nil {
				fmt.Println(err.Error())
				return
			}
			switch validateUser(request.Form.Get("login"), request.Form.Get("pass")) {
			case true:
				answer = "Logged in"
			default:
				answer = "Wrong password/username"
			}
		case jsonFormType:
			mymap := make(map[string]string)
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			err = json.Unmarshal(body, &mymap)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			switch validateUser(mymap["login"], mymap["pass"]) {
			case true:
				answer = "Logged in"
			default:
				answer = "Wrong password/username"
			}
		default:
			answer = "Wrong content-type!"
		}
		fmt.Fprint(writer, answer)
	}
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/login", login)
	http.HandleFunc("/reg", register)
	http.ListenAndServe(":8080", nil)
}
