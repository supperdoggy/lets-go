package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type enterToken struct {
	token     bson.ObjectId
	limited   bool
	savedTime int64
}

func (t *enterToken) expired(minutes int64) (result bool) {
	if t.limited == true {
		result = !(((time.Now().Unix() - t.savedTime) / 60) > minutes)
	}
	return false
}
