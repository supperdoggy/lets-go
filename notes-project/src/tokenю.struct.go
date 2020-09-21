package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"time"
)

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

func validateEntryToken(s *string) bool {
	t, ok := tokenCache[*s]
	if ok {
		if !t.expired(5) {
			return true
		}
		delete(tokenCache, *s)
		return false
	}
	return false
}
