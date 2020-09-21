package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type messageStruct struct {
	Id       bson.ObjectId `bson:"_id" form:"_id" json:"_id"`
	Author   string        `bson:"author" form:"author" json:"author"`
	Msg      string        `bson:"msg" form:"msg" json:"msg"`
	DateTime time.Time     `bson:"date" form:"date" json:"date"`
}

func (m *messageStruct) string() string {
	return m.Author + ": " + m.Msg + "\n"
}
