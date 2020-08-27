package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const jsonContentType = "application/json"

// reading plain text answer got from server
func readPlainTextResponse(response *http.Response) (string, error) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(body), nil
}

func postRequestJson(u string, val map[string]string) (answer string, err error) {
	// encoding values
	out, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(out)
	// sending POST request to server
	resp, err := http.Post(u, jsonContentType, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// returning answer
	return readPlainTextResponse(resp)
}

func getNewToken(username, password string) (string, error) {
	userData := map[string]string{
		"username": username,
		"pass":     password,
	}
	answer, err := postRequestJson("http://localhost:8080/getToken", userData)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	if answer == "Wrong username/password" {
		return "", fmt.Errorf("wrong username/password")
	}
	return answer, nil
}

func ping(userToken string) (err error) {
	userData := map[string]string{
		"token": userToken,
	}
	answer, err := postRequestJson("http://localhost:8080/ping", userData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if answer == "Wrong token!" {
		return fmt.Errorf(answer)
	}
	fmt.Println(answer)
	return
}

func main() {
	userToken, err := getNewToken("admin", "admin")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Token = ", userToken)
		err = ping(userToken)
		if err != nil && err.Error() == "Wrong token!" {
			userToken, err = getNewToken("admin", "admin")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		time.Sleep(1 * time.Minute)
	}
}
