package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

const (
	mongoUrl = "mongodb://127.0.0.1:27017/"
	dbName = "gin-login"
	usersSessionName = "users"
	salt = "meme123meme123"
)

// sha256 hashing algorithm
func getSha256(text string) string {
	hashser := sha256.New()
	hashser.Write([]byte(text + salt))
	return hex.EncodeToString(hashser.Sum(nil))
}

func mainPage(c *gin.Context){
	if s, err := c.Cookie("login"); s == "true" && err == nil{
		c.String(http.StatusOK, "Logged in!")
		return
	}
	c.String(http.StatusOK, "You need to login")
	return
}

func loginPage(c *gin.Context){
	c.File("temples/login.html")
}

func registerPage(c *gin.Context){
	c.File("temples/register.html")
}

func usernameIsTaken(users *mgo.Collection, username string) (result bool, err error){
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

func register(c *gin.Context){
	username := c.PostForm("login")
	password := c.PostForm("pass")
	if username == "" && password == ""{
		c.Redirect(http.StatusMovedPermanently, "/register")
		return
	}
	u := User{
		Id:       bson.NewObjectId(),
		Username: username,
		Password: getSha256(password),
	}
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if session==nil{
		err = fmt.Errorf("session is nil")
		fmt.Println(err.Error())
		return
	}
	defer session.Close()
	users := session.DB(dbName).C(usersSessionName)
	taken, err :=  usernameIsTaken(users, username)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	if taken{
		c.Redirect(http.StatusMovedPermanently, "/register")
		return
	}
	err = users.Insert(u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func validateUser(username, password string) bool {
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	users := session.DB(dbName).C(usersSessionName)
	foundUsers := []User{}
	err = users.Find(bson.M{"username": username, "password": password}).All(&foundUsers)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if len(foundUsers) == 1 {
		return true
	}
	return false
}

func login(c *gin.Context){
	// getting form data
	username := c.PostForm("login")
	password := c.PostForm("pass")
	if username == "" && password == ""{
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	if validateUser(username, getSha256(password)){
		c.SetCookie("login", "true", 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func main(){
	fmt.Println("Starting server...")
	r := gin.Default()
	// login
	r.GET("/login", loginPage)
	r.POST("/login", login)
	// register
	r.GET("/register", registerPage)
	r.POST("/register", register)
	// main path
	r.GET("/", mainPage)
	if err := r.Run(); err != nil{
		fmt.Println(err.Error())
		return
	}
}
