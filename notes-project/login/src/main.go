package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func getBcrypt(text string) string {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(text), 4)
	if err != nil {
		panic(err.Error())
	}
	return string(hashedPass)
}

func main() {
	fmt.Println("Starting server...")
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/login", login)
		api.POST("/register", register)
	}

	if err := r.Run(":2283"); err != nil {
		fmt.Println(err.Error())
		return
	}
}
