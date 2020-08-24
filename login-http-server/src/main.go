package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

var users = make(map[string]string)

func getSha256(text string) string{
	hashser := sha256.New()
	hashser.Write([]byte(text))
	return hex.EncodeToString(hashser.Sum(nil))
}

func register(writer http.ResponseWriter, request *http.Request)  {

}

func login(writer http.ResponseWriter, request *http.Request){
	switch request.Method {
	case "GET":
		fmt.Println("Got get request")
		http.ServeFile(writer, request, "login.html")
	case "POST":
		fmt.Println("Got post request")
		if err:=request.ParseForm();err!=nil{
			fmt.Println(err.Error())
			return
		}
		username := request.Form.Get("login")
		if pass, ok := users[username];ok{
			if pass == getSha256(request.Form.Get("pass")){
				fmt.Fprint(writer, "Logged in!")
				return
			}
			fmt.Fprint(writer, "Wrong password")
		}else{
		pass := getSha256(request.Form.Get("pass"))
		users[username] = pass
		fmt.Fprint(writer, "Registered")
	}
	}
}

func main(){
	fmt.Println("Starting server...")
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
