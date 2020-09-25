package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func updateNote(c *gin.Context)  {
	//checkLogin(c)

	notesCollection, err := getMongoSession(dbName, notesSessionName)
	if err != nil {
		panic(err.Error())
		return
	}

	id := c.Param("id")
	Title := c.PostForm("Title")
	Text := c.PostForm("Text")

	if id == ""{
		fmt.Println(err.Error())
		c.Redirect(308, "/")
		return
	}

	selector := bson.M{"publicId":id}
	update := bson.M{"$set":bson.M{"title":Title, "text":Text}}

	err = notesCollection.Update(selector, update)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func newNote(c *gin.Context){
	//checkLogin(c)

	notesSession, err := getMongoSession(dbName, notesSessionName)
	if err != nil{
		panic(err.Error())
		return
	}

	// here ill take username from token
	//
	//t, err := c.Cookie("t")
	//if err != nil{
	//	c.Redirect(308, "auth/login")
	//	return
	//}
	//
	//token, err := findTokenStructInMap(t)
	//if err != nil {
	//	fmt.Println("Didn`t find token")
	//	c.Redirect(308, "auth/login")
	//	return
	//}
	//

	Title := c.PostForm("Title")
	Text := c.PostForm("Text")
	Username := c.PostForm("Username")
	fmt.Println(Title, Text, Username)
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
	if err != nil{
		panic(err.Error())
		return
	}

}

func shareNote(c *gin.Context){
	//noteSession, err := getMongoSession(dbName, notesSessionName)
	//if err != nil {
	//	panic(err.Error())
	//	return
	//}
	//username := c.Param("username")

}
