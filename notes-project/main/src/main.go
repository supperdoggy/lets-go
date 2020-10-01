package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
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
	checkLogin(c)
	// getting token struct
	t, err := c.Cookie("t")
	if err != nil {
		c.Redirect(308, "/auth/login")
		return
	}
	token, _ := tokenCache[t]
	notesSession, _ := getMongoSession(dbName, notesSessionName)

	id := c.Param("id")

	var result Note
	err = notesSession.Find(bson.M{"publicId": id}).One(&result)
	if err != nil {
		panic(err.Error())
		return
	}

	c.HTML(200, "comment.html", bson.M{"token": token, "note": result, "id": id})
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
		m.GET("/note/:id", notePage)
	}
	//api := r.Group("/api")
	//{
	//	api.POST("/newNote", n      ewNote)
	//	api.POST("/updateNote", updateNote)
	//	api.POST("/share", shareNote)
	//	api.GET("/test", func(c *gin.Context) {
	//		c.HTML(200, "postRequestAjax.html", bson.M{})
	//	})
	//	api.POST("/test", func(c *gin.Context){
	//		var data = make(map[string]interface{})
	//		d := c.PostForm("request")
	//		if d == "ping"{
	//			data["response"] = "pong"
	//		}else if d == "pong"{
	//			data["response"] = "ping"
	//		}
	//		fmt.Println(data, d)
	//		c.JSON(200, data)
	//	})
	//}

	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
