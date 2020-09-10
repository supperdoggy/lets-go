package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	fmt.Println("GET")
}

func post(c *gin.Context) {
	fmt.Println("POST")
	var res map[string]interface{}
	if err := c.Bind(&res);err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
	c.JSON(200, gin.H{"res":res})
}

func main() {
	r := gin.Default()
	r.GET("/ping", get)
	r.POST("/ping", post)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
