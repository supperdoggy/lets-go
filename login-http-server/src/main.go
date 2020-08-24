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

// map with user data, username and hashed password
var users = make(map[string]string)

// sha256 hashing algorithm
func getSha256(text string) string {
	hashser := sha256.New()
	hashser.Write([]byte(text))
	return hex.EncodeToString(hashser.Sum(nil))
}

// path for user registration
func register(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		fmt.Println("Got get request")
		http.ServeFile(writer, request, "reg.html")
	case "POST":
		fmt.Println("Got post request")
		// parsing form
		if err := request.ParseForm(); err != nil {
			fmt.Println(err.Error())
			return
		}
		username := request.Form.Get("login")
		hashedPassword := getSha256(request.Form.Get("pass"))
		// creating user
		users[username] = hashedPassword
		fmt.Fprint(writer, "Registered")
	}
}

// checks if user is in map
func validateUser(username, password string) bool {
	if savedPassword, ok := users[username]; ok {
		if savedPassword == password {
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
			// checking if user is in map
			switch validateUser(request.Form.Get("login"), getSha256(request.Form.Get("pass"))) {
			case true:
				answer = "Logged in"
			default:
				answer = "Wrong password/username"
			}
		case jsonFormType:
			// getting user json request
			jsonRequest := make(map[string]string)
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			err = json.Unmarshal(body, &jsonRequest)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			// checking if user is in map
			switch validateUser(jsonRequest["login"], getSha256(jsonRequest["pass"])) {
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
	// login path
	http.HandleFunc("/login", login)
	// register path
	http.HandleFunc("/reg", register)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}
