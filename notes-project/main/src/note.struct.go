package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Note struct {
	Id       bson.ObjectId          `bson:"_id" form:"_id" json:"_id"`
	PublicId string                 `bson:"publicId" form:"publicId" json:"publicId"`
	Title    string                 `bson:"title" form:"title" json:"title"`
	Text     string                 `bson:"text" form:"text" json:"text"`
	Owner    string                 `bson:"owner" form:"owner" json:"owner"`
	Created  time.Time              `bson:"created" form:"created" json:"created"`
	Shared   bool                   `bson:"shared" form:"shared" json:"shared"`
	Users    map[string]Permissions `bson:"users" form:"users" json:"users"`
}

func (n *Note) addNewUser(userId string, p Permissions) error {
	if !n.Shared {
		return fmt.Errorf("note is not shared")
	}
	if _, ok := n.Users[userId]; ok {
		return fmt.Errorf("user already in map")
	}
	n.Users[userId] = p
	return nil
}

func (n *Note) deleteUser(userId string) error {
	if !n.Shared {
		return fmt.Errorf("note is not shared")
	}
	delete(n.Users, userId)
	return nil
}

func (n *Note) changePermission(userId string, p Permissions) error {
	if !n.Shared {
		return fmt.Errorf("note is not shared")
	}
	if _, ok := n.Users[userId]; !ok {
		return fmt.Errorf("user already in map")
	}
	n.Users[userId] = p
	return nil
}
