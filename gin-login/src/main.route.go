package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func mainPage(c *gin.Context) {
	checkLogin(c)
	s, err := mgo.Dial(mongoUrl)
	if err != nil {
		c.String(200, err.Error())
		return
	}
	defer s.Close()
	msgCollection := s.DB(dbName).C(messageSessionName)
	var messages []messageStruct
	iter := msgCollection.Find(nil).Limit(20).Iter()
	err = iter.All(&messages)
	if err != nil {
		panic(err.Error())
		return
	}
	username, err := c.Cookie("username")
	if err != nil {
		c.Redirect(http.StatusPermanentRedirect, "/auth/login")
		return
	}
	c.HTML(200, "main.html", gin.H{"messages": messages, "username": username})
	return
}

func mainNewMessage(c *gin.Context) {
	checkLogin(c)
	username, err := c.Cookie("username")
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusPermanentRedirect, "/auth/login")
		return
	}
	msg := c.PostForm("msg")
	if username != "" && msg != "" {
		newMessage(username, msg)
	}
	c.Redirect(http.StatusPermanentRedirect, "/")
	return
}

func newMessage(username, msg string) {
	s, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	defer s.Close()
	messages := s.DB(dbName).C(messageSessionName)

	m := messageStruct{
		Id:       bson.NewObjectId(),
		Author:   username,
		Msg:      msg,
		DateTime: time.Now(),
	}
	err = messages.Insert(m)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	return
}
