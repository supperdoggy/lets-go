package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func mainPage(c *gin.Context) {
	checkLogin(c)
	t, err := c.Cookie("t")
	if err != nil {
		// if we get an error returning user to login page
		c.Redirect(http.StatusPermanentRedirect, "auth/login")
		return
	}
	token, ok := tokenCache[t]
	if !ok {
		c.Redirect(http.StatusPermanentRedirect, "auth/login")
		return
	}
	c.HTML(200, "index.html", gin.H{"token":token})
	return
}
