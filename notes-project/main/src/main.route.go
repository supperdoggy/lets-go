package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func processAnswer(i []interface{}) []map[string]interface{} {
	result := make([]map[string]interface{}, 1)
	for _, v := range i {
		m := v.(map[string]interface{})
		result = append(result, m)
	}
	return result
}

func mainPage(c *gin.Context) {
	checkLogin(c)
	t, err := c.Cookie("t")
	if err != nil {
		// if we get an error returning user to login page
		c.Redirect(http.StatusPermanentRedirect, "auth/login")
		return
	}
	token, err := findTokenStructInMap(t)
	if err != nil {
		c.Redirect(http.StatusPermanentRedirect, "auth/login")
		return
	}
	resp, err := http.PostForm("http://localhost:2020/api/getNotes", url.Values{
		"username": {token},
	})
	var notes map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&notes)
	if err != nil {
		panic(err.Error())
		return
	}

	answer, _ := notes["answer"].(map[string]interface{})
	var ownInterface []interface{}
	var sharedInterface []interface{}
	if answer["ownedNotes"] != nil {
		ownInterface = answer["ownedNotes"].([]interface{})
	}
	if answer["sharedNotes"] != nil {
		sharedInterface = answer["sharedNotes"].([]interface{})
	}

	own := processAnswer(ownInterface)[1:]
	shared := processAnswer(sharedInterface)[1:]
	c.HTML(200, "index1.html", gin.H{"token": token, "own": own, "shared": shared})
	return
}
