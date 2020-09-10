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
	usersCollectionName = "users1"
)

func usernameIsTaken(users *mgo.Collection, username string) (result bool, err error){
	foundUsers := []User{}
	err = users.Find(bson.M{"username": username}).All(&foundUsers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// if user already in base
	if len(foundUsers) > 0 {
		result = true
		return
	}
	return
}

func insertIntoCollection(collection *mgo.Collection, data interface{}) (err error){
	err = collection.Insert(data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

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
		// if user already in base
		isTaken, err := usernameIsTaken(users, username)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if isTaken{
			_, err = fmt.Fprintln(writer, "Username is taken")
			if err !=nil{
				fmt.Println(err.Error())
				return
			}
			return
		}
		err = users.EnsureIndex(mgo.Index{
			Key:              []string{"Username"},
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// writing user into db
		err = insertIntoCollection(users, u)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		return
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
		_, err := fmt.Fprint(writer, answer)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
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
