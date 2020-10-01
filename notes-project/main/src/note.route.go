package main

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"gopkg.in/mgo.v2/bson"
//	"net/http"
//	"net/url"
//	"strconv"
//	"time"
//)
//
//func newNote(c *gin.Context) {
//	//checkLogin(c)
//
//	notesSession, err := getMongoSession(dbName, notesSessionName)
//	if err != nil {
//		panic(err.Error())
//		return
//	}
//
//	// here ill take username from token
//	//
//	//t, err := c.Cookie("t")
//	//if err != nil{
//	//	c.Redirect(308, "auth/login")
//	//	return
//	//}
//	//
//	//token, err := findTokenStructInMap(t)
//	//if err != nil {
//	//	fmt.Println("Didn`t find token")
//	//	c.Redirect(308, "auth/login")
//	//	return
//	//}
//	//
//
//	Title := c.PostForm("Title")
//	Text := c.PostForm("Text")
//	Username := c.PostForm("Username")
//	fmt.Println(Title, Text, Username)
//	note := Note{
//		Id:       bson.NewObjectId(),
//		PublicId: bson.NewObjectId().String(),
//		Title:    Title,
//		Text:     Text,
//		Owner:    Username, // token
//		Created:  time.Now(),
//		Shared:   false,
//		Users:    nil,
//	}
//	err = notesSession.Insert(note)
//	if err != nil {
//		panic(err.Error())
//		return
//	}
//
//}
//
//func shareNote(c *gin.Context) {
//
//	owner := c.PostForm("Owner")
//	username := c.PostForm("Username") // username of user we want to share note with
//	id := c.PostForm("Id") // public id
//	canRedact, _ := strconv.ParseBool(c.PostForm("CanRedact"))
//	canAddNewUsers, _ := strconv.ParseBool(c.PostForm("CanAddNewUsers"))
//
//	err := http.PostForm("")
//
//}
