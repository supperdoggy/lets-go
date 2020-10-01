package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type enterToken struct {
	Token     string `bson:"token" json:"token" form:"token"`
	Limited   bool          `bson:"limited" json:"limited" form:"limited"`
	SavedTime int64         `bson:"savedTime" json:"saved_time" form:"savedTime"`
	Username  string        `bson:"username" json:"username"`
}

func (t *enterToken) expired(minutes int64) (result bool) {
	if t.Limited == true {
		result = !(((time.Now().Unix() - t.SavedTime) / 60) > minutes)
	}
	return false
}

func createNewToken(limited bool, username string) enterToken {
	t := enterToken{
		Token:    strconv.FormatInt(time.Now().UnixNano(), 10),
		Username: username,
	}
	if limited {
		t.Limited = true
		t.SavedTime = time.Now().Unix()
	}
	return t
}

func validateEntryToken(s *string) bool {
	t, ok := tokenCache[*s]
	if ok {
		if !t.expired(30) { // expiration time of token is set to 30 minutes
			return true
		}
		delete(tokenCache, *s)
		return false
	}
	return false
}

func deleteCookieFromMap(c *gin.Context) {
	t := c.PostForm("t")
	if t != "" {
		return
	}else{
		delete(tokenCache, t)
		return
	}
}

func findTokenStructInMap(t string) (enterToken, error) {
	if token, ok := tokenCache[t]; ok {
		return token, nil
	}
	return enterToken{}, fmt.Errorf("token is not valid")
}