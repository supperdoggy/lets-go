package main

import (
	"github.com/gin-gonic/gin"
)

func mainPage(c *gin.Context) {
	checkLogin(c)

	c.String(200, "Logged in!")
	return
}
