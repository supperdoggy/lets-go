package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func mainPage(c *gin.Context) {
	checkLogin(c)
	c.HTML(200, "main.html", gin.H{})
	return
}

func mainNewMessage(c *gin.Context) {
	checkLogin(c)
	username := c.PostForm("username")
	msg := c.PostForm("msg")
	if username != "" && msg != "" {
		newMessage(username, msg)
	}
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
		id:       bson.NewObjectId(),
		author:   username,
		msg:      msg,
		dateTime: time.Now(),
	}
	err = messages.Insert(m)
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	return
}
