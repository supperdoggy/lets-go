package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"strings"
	"time"
)

func login(c *gin.Context) {
	response := map[string]interface{}{
		"ok": false,
		"error": "",
		"answer": false,
	}
	usersCollection, err := getMongoSession(dbName, usersSessionName)
	if err != nil {
		response["error"] = "Error connecting to mongodb"
		c.JSON(200, response)
		return
	}

	username := c.PostForm("login")
	password := c.PostForm("pass")
	if username == "" || password == "" {
		response["error"] = "not all inputs are filled"
		c.JSON(200, response)
		return
	}
	if validateUser(username, password, usersCollection) {
		response["ok"] = true
		response["answer"] = true
		c.JSON(200, response)
		return
	}
	response["ok"] = true
	c.JSON(200, response)
	return
}

func validateUser(username, password string, users *mgo.Collection) bool {
	foundUsers := []User{}
	err := users.Find(bson.M{"username": strings.ToLower(username)}).All(&foundUsers)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(foundUsers[0].Password), []byte(password)); err == nil {
		return true
	}
	return false
}

func register(c *gin.Context) {
	response := map[string]interface{}{
		"ok": false,
		"error": "",
		"answer": false,
	}

	username := c.PostForm("login")
	password := c.PostForm("pass")
	fmt.Println(username, password)
	if username == "" || password == "" {
		c.JSON(200, map[string]interface{}{})
		return
	}
	u := User{
		Id:       bson.NewObjectId(),
		UniqueId: strconv.FormatInt(time.Now().UnixNano(), 10),
		Username: username,
		Password: getBcrypt(password),
		Created:  time.Time{},
	}

	usersCollection, err := getMongoSession(dbName, usersSessionName)
	if err != nil {
		fmt.Println(err.Error())
		response["error"] = "Error connecting to mongodb"
		c.JSON(200, response)
	}
	taken, err := usernameIsTaken(usersCollection, username)
	if err != nil {
		response["error"] = "cant find user in db"
		c.JSON(200, response)
		return
	}
	if taken {
		response["error"] = "Username is already taken!"
		c.JSON(200, response)
		return
	}
	err = usersCollection.Insert(u)
	if err != nil {
		response["error"] = "error inserting new user"
		c.JSON(200, response)
		return
	}
	response["ok"] = true
	response["answer"] = true
	c.JSON(200, response)
	return
}

func usernameIsTaken(users *mgo.Collection, username string) (result bool, err error) {
	foundUsers := []User{}
	err = users.Find(bson.M{"username": strings.ToLower(username)}).All(&foundUsers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// if user already in base
	if len(foundUsers) > 0 {
		result = true
		return
	}
	return
}

func validateToken(c *gin.Context){
	t := c.PostForm("t")
	response := map[string]interface{}{
		"ok":true,
		"error":"",
		"answer":false,
	}
	if validateEntryToken(&t){
		response["answer"] = true
	}else{
		response["answer"] = false
	}
	c.JSON(200, response)
	return
}
