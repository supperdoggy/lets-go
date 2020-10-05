package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func updateNote(c *gin.Context) {
	id := c.PostForm("id")
	Text := c.PostForm("Text")
	Title := c.PostForm("Title")

	if id == "" {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  "id is empty",
			"answer": false,
		})
		return
	}
	selector := bson.M{"publicId": id}
	update := bson.M{"$set": bson.M{"text": Text, "title": Title}}

	err = notesSession.Update(selector, update)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"ok":     true,
		"error":  "",
		"answer": true,
	})
	return
}

func newNote(c *gin.Context) {
	Title := c.PostForm("Title")
	Text := c.PostForm("Text")
	Username := c.PostForm("Username")
	if Title == "" || Username == "" {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  "not all fields are filled",
			"answer": false,
		})
		return
	}
	note := Note{
		Id:       bson.NewObjectId(),
		PublicId: strconv.FormatInt(time.Now().UnixNano(), 10), // just taking current nanosecs in unix
		Title:    Title,
		Text:     Text,
		Owner:    Username,
		Created:  time.Now(),
		Shared:   false,
		Users:    nil,
	}
	err = notesSession.Insert(note)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"ok":     true,
		"error":  "",
		"answer": true,
	})
	return
}

func getNote(c *gin.Context) {
	response := map[string]interface{}{
		"ok":     false,
		"error":  "",
		"answer": nil,
	}
	id := c.PostForm("id")
	t := c.PostForm("t")
	var username string

	var result Note
	err = notesSession.Find(bson.M{"publicId": id}).One(&result)
	if err != nil {
		response["error"] = err.Error()
		c.JSON(200, response)
		return
	}

	// getting username
	data := make(map[string]interface{})
	resp, err := http.PostForm("http://localhost:2283/api/getTokenStruct", url.Values{"t": {t}})
	if err != nil {
		response["error"] = err.Error()
		c.JSON(200, response)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		response["error"] = err.Error()
		c.JSON(200, response)
		return
	}
	// checking if user allowed to take get note
	if data["ok"].(bool) {
		username = data["answer"].(string)
	} else {
		response["error"] = "wrong request"
		c.JSON(200, response)
		return
	}
	if result.Owner == username {
		response["ok"] = true
		response["answer"] = result
		c.JSON(200, response)
		return
	} else if result.Shared {
		if _, ok := result.Users[username]; ok {
			response["ok"] = true
			response["answer"] = result
			c.JSON(200, response)
			return
		}
	} else {
		response["error"] = "not allowed"
		c.JSON(200, response)
		return
	}
}

func shareNote(c *gin.Context) {
	owner := c.PostForm("Owner")
	username := c.PostForm("Username") // username of user we want to share note with
	id := c.PostForm("Id")             // public id
	canRedact := true                  // default value
	canAddNewUsers := true             // default value
	var note Note

	err = notesSession.Find(bson.M{"publicId": id}).One(&note)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}

	note.shareNote()
	err = note.addNewUser(username, Permissions{
		CanRedact:      canRedact,
		CanAddNewUsers: canAddNewUsers,
	})
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}
	selector := bson.M{
		"publicId": id,
		"owner":    owner,
	}
	update := bson.M{
		"$set": bson.M{
			"shared": true,
			"users":  note.Users,
		},
	}
	err = notesSession.Update(selector, update)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"ok":     true,
		"error":  "",
		"answer": true,
	})
	return
}

func sendNotes(c *gin.Context) {
	username := c.PostForm("username")
	var ownedNotes []Note
	err = notesSession.Find(bson.M{"owner": username}).All(&ownedNotes)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": "",
		})
		return
	}
	var sharedNotes []Note
	err = notesSession.Find(bson.M{"shared": true, "users." + username: bson.M{"$exists": true}}).All(&sharedNotes)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"ownedNotes":  ownedNotes,
		"sharedNotes": sharedNotes,
	})
	return
}

func deleteNote(c *gin.Context) {
	response := map[string]interface{}{
		"ok":    false,
		"error": "",
	}
	id := c.PostForm("id")

	err = notesSession.Remove(bson.M{"publicId": id})
	if err != nil {
		response["error"] = err.Error()
		c.JSON(200, response)
		return
	}
	response["ok"] = true
	c.JSON(200, response)
	return
}
