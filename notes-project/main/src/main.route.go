package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func mainPage(c *gin.Context) {
	checkLogin(c)
	fmt.Println("token is good")
	t, err := c.Cookie("t")
	if err != nil {
		// if we get an error returning user to login page
		c.Redirect(http.StatusPermanentRedirect, "auth/login")
		return
	}
	token, err := findTokenStructInMap(t)
	if err!=nil {
		c.Redirect(http.StatusPermanentRedirect, "auth/login")
		return
	}
	resp, err := http.PostForm("http://localhost:2020/api/getNotes", url.Values{
		"username":{token},
	})
	var notes map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&notes)
	if err != nil{
		panic(err.Error())
		return
	}

	// how to work with data
	//answer, _ := notes["answer"].(map[string]interface{})
	//own := answer["ownedNotes"].([]interface{})
	//shared := answer["sharedNotes"]
	//fmt.Println(own[0].(map[string]interface{})["text"])
	//fmt.Println(shared)

	c.HTML(200, "index.html", gin.H{"token": token, "notes":notes["answer"]})
	return
}
