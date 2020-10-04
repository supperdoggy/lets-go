package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
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
	if cookie != "" {
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
	if cookie != "" {
		data["error"] = cookie
	}

	c.HTML(200, "register.html", data)
}

func notePage(c *gin.Context) {
	id := c.Param("id")
	checkLogin(c)
	// getting token struct
	t, err := c.Cookie("t")
	if err != nil {
		c.Redirect(308, "/auth/login")
		return
	}
	token, err := findTokenStructInMap(t)
	if err != nil {
		c.Redirect(308, "/auth/login")
		return
	}
	result := make(map[string]interface{})
	resp, err := http.PostForm("http://localhost:2020/api/getNote", url.Values{"id":{id}, "t":{t}})
	if err != nil{
		c.Redirect(308, "/")
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil{
		c.Redirect(308, "/")
		return
	}
	if result["ok"].(bool){
		note := result["answer"].(map[string]interface{})
		c.HTML(200, "comment.html", bson.M{"token": token, "note": note, "id": id})
		return
	}else{
		c.Redirect(308, "/")
	}
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
		m.Any("/note/:id", notePage)
	}

	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}