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
	cookie, err := c.Cookie("error")
	if err != nil {
		cookie = ""
	}
	data := gin.H{}
	if cookie != ""{
		data["error"] = cookie
	}

	c.HTML(200, "login.html", data)
}

func registerPage(c *gin.Context) {
	cookie, err := c.Cookie("error")
	if err != nil {
		cookie = ""
	}
	data := gin.H{}
	if cookie != ""{
		data["error"] = cookie
	}

	c.HTML(200, "register.html", data)
}

func main() {
	fmt.Println("Starting server...")
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

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
	}

	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
