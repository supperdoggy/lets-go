package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

func login(c *gin.Context) {
	usersCollection, err := getMongoSession(dbName, usersSessionName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, map[string]interface{}{"ok": false, "error": "Error connecting to mongodb", "answer": false})
		return
	}

	username := c.PostForm("login")
	password := c.PostForm("pass")
	if username == "" || password == "" {
		c.JSON(200, map[string]interface{}{"ok": false, "error": "not all inputs are filled", "answer": false})
		return
	}
	if validateUser(username, password, usersCollection) {
		c.JSON(200, map[string]interface{}{"ok": true, "error": "", "answer": true})
		return
	}
	c.JSON(200, map[string]interface{}{"ok": true, "error": "", "answer": false})
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
	username := c.PostForm("login")
	password := c.PostForm("pass")
	fmt.Println(username, password)
	if username == "" || password == "" {
		c.JSON(200, map[string]interface{}{"ok": false, "error": "not all inputs are filled", "answer": false})
		return
	}
	u := User{
		Id:       bson.NewObjectId(),
		UniqueId: bson.NewObjectId(),
		Name:     "name",
		Username: username,
		Password: getBcrypt(password),
		Created:  time.Time{},
	}
	usersCollection, err := getMongoSession(dbName, usersSessionName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, map[string]interface{}{"ok": false, "error": "Error connecting to mongodb"})
	}
	taken, err := usernameIsTaken(usersCollection, username)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, map[string]interface{}{"ok": false, "error": "cant find user in db"})
		return
	}
	if taken {
		c.JSON(200, map[string]interface{}{"ok": false, "error": "Username is already taken!"})
		return
	}
	err = usersCollection.Insert(u)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, map[string]interface{}{"ok": false, "error": "error inserting new user"})
		return
	}
	c.JSON(200, map[string]interface{}{"ok": true, "error": "", "answer": true})
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
