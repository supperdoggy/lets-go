package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type messageStruct struct {
	id       bson.ObjectId `bson:"_id" form:"_id" json:"_id"`
	author   string        `bson:"author" form:"author" json:"author"`
	msg      string        `bson:"msg" form:"msg" json:"msg"`
	dateTime time.Time     `bson:"date" form:"date" json:"date"`
}

func (m *messageStruct) string() string {
	return m.author + ": " + m.msg + "\n"
}
