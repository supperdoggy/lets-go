package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	mongoUrl           = "mongodb://127.0.0.1:27017/"
	dbName             = "gin-login"
	usersSessionName   = "users"
	messageSessionName = "messages"
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

func loginPage(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func registerPage(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}

func main() {
	fmt.Println("Starting server...")
	r := gin.Default()
	r.LoadHTMLGlob("temples/*")

	auth := r.Group("/auth")
	{
		// login
		auth.GET("/login", loginPage)
		auth.POST("/login", login)
		// register
		auth.GET("/register", registerPage)
		auth.POST("/register", register)

	}

	r.GET("/", mainPage)
	r.POST("/newMessage", mainNewMessage)
	//// validating login token
	//r.GET("/validate", checkLogin)

	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
