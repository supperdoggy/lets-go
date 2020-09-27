package main

import "gopkg.in/mgo.v2"

func getMongoSession(db, collectionName string) (*mgo.Collection, error) {
	s, err := mgo.Dial(mongoUrl)
	if err != nil {
		panic(err.Error())
		return nil, err
	}
	return s.DB(db).C(collectionName), nil
}
