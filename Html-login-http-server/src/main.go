package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

const (
	htmlFormType        = "application/x-www-form-urlencoded"
	jsonFormType        = "application/json"
	salt                = "memes123memes123"
	mongoUrl            = "mongodb://127.0.0.1:27017/"
	usersDbName         = "httpLoginServer"
	usersCollectionName = "users"
)

// map with user data, username and hashed password
var users = make(map[string]string)

// sha256 hashing algorithm
func getSha256(text string) string {
	hashser := sha256.New()
	hashser.Write([]byte(text + salt))
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
		// getting data
		username := request.Form.Get("login")
		hashedPassword := getSha256(request.Form.Get("pass"))
		// creating user
		u := User{
			Id:       bson.NewObjectId(),
			Username: username,
			Password: hashedPassword,
		}
		// connecting to mongo
		session, err := mgo.Dial(mongoUrl)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if session == nil {
			fmt.Println("nil session")
			return
		}
		defer session.Close()
		users := session.DB(usersDbName).C(usersCollectionName)
		// finding out if username is taken
		foundUsers := []User{}
		err = users.Find(bson.M{"username": username}).All(&foundUsers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// if user already in base
		if len(foundUsers) > 0 {
			_, err = fmt.Fprint(writer, "Username if already taken")
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}
		// writing user into db
		err = users.Insert(u)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		_, err = fmt.Fprint(writer, "Registered")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

// checks if user is in db
func validateUser(username, password string) bool {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	users := session.DB(usersDbName).C(usersCollectionName)
	foundUsers := []User{}
	err = users.Find(bson.M{"username": username, "password": password}).All(&foundUsers)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if len(foundUsers) == 1 {
		return true
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
	// first user
	users["admin"] = "admin"
	//
	fmt.Println("Starting server...")
	// login path
	http.HandleFunc("/login", login)
	// register path
	http.HandleFunc("/reg", register)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}
