package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type enterToken struct {
	Token     bson.ObjectId `bson:"token" json:"token" form:"token"`
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
		Token:    bson.NewObjectId(),
		Username: username,
	}
	if limited {
		t.Limited = true
		t.SavedTime = time.Now().Unix()
	}
	return t
}

func createNewTokenCookie(c *gin.Context, username string) {
	t := createNewToken(true, username) // cookie is limited
	c.SetCookie("t", t.Token.String(), 400, "/", "localhost", false, true)
	tokenCache[t.Token.String()] = t
}

func validateEntryToken(s *string) bool {
	t, ok := tokenCache[*s]
	if ok {
		if !t.expired(5) { // expiration time of token is set to 5 minutes
			return true
		}
		delete(tokenCache, *s)
		return false
	}
	return false
}

func deleteCookieFromMap(c *gin.Context) error {
	t, err := c.Cookie("t")
	if err != nil {
		return err
	}
	delete(tokenCache, t)
	return nil
}

func findTokenStructInMap(t string) (enterToken, error){
	if token, ok := tokenCache[t]; ok{
		return token, nil
	}
	return enterToken{}, fmt.Errorf("token is not valid")
}
