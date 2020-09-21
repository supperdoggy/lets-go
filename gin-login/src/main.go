package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	r.Static("/static", "./static")
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

	m := r.Group("/")
	{
		m.POST("/", mainPage)
		m.GET("/", mainPage)
		m.POST("/newMessage", mainNewMessage)
	}
	//// validating login token
	//r.GET("/validate", checkLogin)

	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
