package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

const (
	mongoUrl         = "mongodb://127.0.0.1:27017/"
	dbName           = "gin-login"
	usersSessionName = "users"
)

var (
	tokenCache = make(map[string]enterToken)
)

func getBcrypt(text string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(text), 4)
	if err != nil {
		panic(err.Error())
	}
	return string(hashedPass)
}

func createNewToken() enterToken {
	t := enterToken{
		token:     bson.NewObjectId(),
		limited:   true,
		savedTime: time.Now().Unix(),
	}
	return t
}

func createNewTokenCookie(c *gin.Context) {
	t := createNewToken()
	c.SetCookie("t", t.token.String(), 400, "/", "localhost", false, true)
	tokenCache[t.token.String()] = t
}

func validateEntryToken(s *string, c *gin.Context) bool {
	t, ok := tokenCache[*s]
	if ok {
		if !t.expired(5){
			return true
		}
		delete(tokenCache, *s)
		return false
	}
	return false
}

func mainPage(c *gin.Context) {
	t, _ := c.Cookie("t")
	if validateEntryToken(&t, c) {
		c.String(http.StatusOK, "Logged in via cookie")
		return
	}
	delete(tokenCache, t)
	c.String(http.StatusOK, "You need to login")
	return
}

func loginPage(c *gin.Context) {
	c.File("temples/login.html")
}

func registerPage(c *gin.Context) {
	c.File("temples/register.html")
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

func register(c *gin.Context) {
	username := c.PostForm("login")
	password := c.PostForm("pass")
	if username == "" || password == "" {
		c.Redirect(http.StatusMovedPermanently, "/register")
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
		c.Redirect(http.StatusMovedPermanently, "/register")
		return
	}
	err = users.Insert(u)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	createNewTokenCookie(c)
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

func login(c *gin.Context) {
	// getting form data
	username := c.PostForm("login")
	password := c.PostForm("pass")
	if username == "" && password == "" {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	if validateUser(username, password) {
		createNewTokenCookie(c)
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func main() {
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
	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
