package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

func updateNote(c *gin.Context) {
	notesCollection, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}
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

	err = notesCollection.Update(selector, update)
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
	//checkLogin(c)

	notesSession, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}

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

func shareNote(c *gin.Context) {
	fmt.Println("Im here!")
	noteSession, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}

	owner := c.PostForm("Owner")
	username := c.PostForm("Username") // username of user we want to share note with
	id := c.PostForm("Id")             // public id
	canRedact := true // default value
	canAddNewUsers := true // default value
	var note Note

	err = noteSession.Find(bson.M{"publicId": id}).One(&note)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}

	note.shareNote() // returns error if note is shared
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}
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
	err = noteSession.Update(selector, update)
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

// NEED TESTING
func sendNotes(c *gin.Context) {
	fmt.Println("im here!")
	notesSession, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":     false,
			"error":  err.Error(),
			"answer": false,
		})
		return
	}
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
		"ok":    true,
		"error": "",
		"answer": map[string]interface{}{
			"ownedNotes":  ownedNotes,
			"sharedNotes": sharedNotes,
		},
	})
	return
}

func deleteNote(c *gin.Context) {
	response := map[string]interface{}{
		"ok":    false,
		"error": "",
	}
	id := c.PostForm("id")

	notesSession, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		response["error"] = err.Error()
		c.JSON(200, response)
		return
	}
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
