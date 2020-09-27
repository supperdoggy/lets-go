package main

type Permissions struct {
	CanRedact      bool `bson:"canRedact" json:"canRedact" form:"canRedact"`
	CanAddNewUsers bool `bson:"canAddNewUsers" json:"canAddNewUsers" form:"canAddNewUsers"`
}
