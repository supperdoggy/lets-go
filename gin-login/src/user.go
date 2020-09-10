package main

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id bson.ObjectId `bson:"_id" form:"_id" json:"_id"`
	Username string `bson:"username" form:"username" json:"username"`
	Password string `bson:"password" form:"password" json:"password"`
}