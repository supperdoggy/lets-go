package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	salt = "sadas123safnmzxm12321asSDaws"
)

var (
	// contains token + timeWhenTokenWasGiven
	usersToken = make(map[string]int64)
	users      = make(map[string]string)
)

func getJsonData(r *http.Request) (result map[string]string, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func ping(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		fmt.Fprint(writer, "Unsupported method")
		return
	}
	if request.Header.Get("content-type") != "application/json" {
		fmt.Fprint(writer, "Unsupported data format")
	}

	jsonData, err := getJsonData(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	token, ok := jsonData["token"]
	if !ok {
		fmt.Fprint(writer, "No token")
		return
	}
	if validateToken(&token) {
		fmt.Fprint(writer, "pong")
		return
	}
	fmt.Fprint(writer, "Wrong token!")
}

// sha256 hashing algorithm
func getSha256(text string) string {
	hashser := sha256.New()
	hashser.Write([]byte(text + salt))
	return hex.EncodeToString(hashser.Sum(nil))
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

// checking if passed more than 5 minutes
func validateToken(token *string) bool {
	if savedTime, ok := usersToken[*token]; ok {
		if (time.Now().Unix()-savedTime)/60 > 5 {
			return false
		} else {
			return true
		}
	}
	return false
}

func randomStringGenerator(l int) string {
	// letters and number to random pick
	const lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	result := make([]byte, l)
	for i := range result {
		result[i] = lettersAndNumbers[rand.Intn(len(lettersAndNumbers))]
	}
	return string(result)
}

// generating token
func generateToken() string {
	token := randomStringGenerator(24)
	usersToken[token] = time.Now().Unix()
	return token
}

func getToken(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		fmt.Fprint(writer, "Unsupported method")
		return
	}
	if request.Header.Get("content-type") != "application/json" {
		fmt.Fprint(writer, "Unsupported data format")
	}

	jsonData, err := getJsonData(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if pass, ok := users[jsonData["username"]]; ok {
		if pass == jsonData["pass"] {
			fmt.Fprint(writer, generateToken())
			return
		}
	}
	fmt.Fprint(writer, "Wrong username/password")

}

func main() {
	fmt.Println("Starting server...")
	users["admin"] = "admin"
	http.HandleFunc("/getToken", getToken)
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
		return
	}
}
