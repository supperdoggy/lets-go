package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

func updateNote(c *gin.Context) {
	notesCollection, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
			"answer":false,
		})
		return
	}

	id := c.PostForm("id")
	Title := c.PostForm("Title")
	Text := c.PostForm("Text")

	if id == "" {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":"id is empty",
			"answer":false,
		})
		return
	}
	selector := bson.M{"publicId": id}
	update := bson.M{"$set": bson.M{"title": Title, "text": Text}}

	err = notesCollection.Update(selector, update)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
			"answer":false,
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"ok":true,
		"error":"",
		"answer":true,
	})
	return
}

func newNote(c *gin.Context) {
	//checkLogin(c)

	notesSession, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
			"answer":false,
		})
		return
	}

	Title := c.PostForm("Title")
	Text := c.PostForm("Text")
	Username := c.PostForm("Username")
	note := Note{
		Id:       bson.NewObjectId(),
		PublicId: bson.NewObjectId().String(),
		Title:    Title,
		Text:     Text,
		Owner:    Username, // token
		Created:  time.Now(),
		Shared:   false,
		Users:    nil,
	}
	err = notesSession.Insert(note)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
			"answer":false,
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"ok":true,
		"error":"",
		"answer":true,
	})
	return
}

func shareNote(c *gin.Context) {
	noteSession, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
			"answer":false,
		})
		return
	}

	owner := c.PostForm("Owner")
	username := c.PostForm("Username") // username of user we want to share note with
	id := c.PostForm("Id") // public id
	canRedact, _ := strconv.ParseBool(c.PostForm("CanRedact"))
	canAddNewUsers, _ := strconv.ParseBool(c.PostForm("CanAddNewUsers"))

	var note Note
	err = noteSession.Find(bson.M{"publicId":id}).One(&note)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
			"answer":false,
		})
		return
	}
	note.Users[username] = Permissions{
		CanRedact:      canRedact,
		CanAddNewUsers: canAddNewUsers,
	}
	selector := bson.M{
		"publicId":id,
		"owner":owner,
	}
	update := bson.M{
		"$set":bson.M{
			"shared":true,
			"users":note.Users,
		},
	}
	err = noteSession.Update(selector, update)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"ok":true,
		"answer":true,
	})
	return
}

// NEED TESTING
func sendNotes(c *gin.Context){
	notesSession, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
			"answer":false,
		})
		return
	}
	username := c.PostForm("username")
	var ownedNotes []Note
	err = notesSession.Find(bson.M{"owner":username}).All(&ownedNotes)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
			"answer":"",
		})
		return
	}
	var sharedNotes []Note
	err = notesSession.Find(bson.M{"$and" : bson.M{"shared": true, "users."+username:bson.M{"$exist": true}}}).All(&sharedNotes)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"ok":false,
			"error":err.Error(),
			"answer":false,
		})
		return
	}
	c.JSON(200, map[string]interface{}{
		"ok":true,
		"error":"",
		"answer":map[string]interface{}{
			"ownedNotes":ownedNotes,
			"sharedNotes":sharedNotes,
		},
	})
	return
}
