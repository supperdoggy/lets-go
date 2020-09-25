package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id       bson.ObjectId `bson:"_id" form:"_id" json:"_id"`
	UniqueId bson.ObjectId `bson:"uniqueId" form:"uniqueId" json:"unique_id"`
	Name     string        `bson:"name" form:"name" json:"name"`
	Username string        `bson:"username" form:"username" json:"username"`
	Password string        `bson:"password" form:"password" json:"password"`
	Created  time.Time     `bson:"created" form:"created" json:"created"`
}
