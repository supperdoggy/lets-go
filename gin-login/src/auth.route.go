package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func login(c *gin.Context) {
	// getting form data
	username := c.PostForm("login")
	password := c.PostForm("pass")
	if username == "" || password == "" {
		c.Redirect(http.StatusMovedPermanently, "/auth/login")
		return
	}
	if validateUser(username, password) {
		createNewTokenCookie(c)
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/auth/login")
}

func validateUser(username, password string) bool {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	users := session.DB(dbName).C(usersSessionName)
	foundUsers := []User{}
	err = users.Find(bson.M{"username": username}).All(&foundUsers)
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
	if username == "" || password == "" {
		c.Redirect(http.StatusMovedPermanently, "/auth/register")
		return
	}
	u := User{
		Id:       bson.NewObjectId(),
		Username: username,
		Password: getBcrypt(password),
	}
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if session == nil {
		err = fmt.Errorf("session is nil")
		fmt.Println(err.Error())
		return
	}
	defer session.Close()
	users := session.DB(dbName).C(usersSessionName)
	taken, err := usernameIsTaken(users, username)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if taken {
		c.Redirect(http.StatusMovedPermanently, "/auth/register")
		return
	}
	err = users.Insert(u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	createNewTokenCookie(c)
	c.Redirect(http.StatusMovedPermanently, "/auth/login")
}

func usernameIsTaken(users *mgo.Collection, username string) (result bool, err error) {
	foundUsers := []User{}
	err = users.Find(bson.M{"username": username}).All(&foundUsers)
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

func checkLogin(c *gin.Context) {
	t, err := c.Cookie("t")
	if err != nil {
		c.Redirect(http.StatusPermanentRedirect, "/auth/login")
		return
	}
	if !validateEntryToken(&t) {
		c.Redirect(http.StatusPermanentRedirect, "/auth/login")
		return
	}
}
