package main


import(
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type student struct {
	Id bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Age uint `bson:"age"`
	Subjects []string `bson:"subjects"`
}

func main(){
	// connecting to mongo db
	session, err := mgo.Dial("mongodb://127.0.0.1:27017/")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if session == nil{
		fmt.Println("session is nil")
		return
	}
	defer session.Close()
	// connecting to db and collection
	students := session.DB("test").C("students")

	// INSERT
	//st := student{
	//	Id:       bson.NewObjectId(),
	//	Name:     "John",
	//	Age:      44,
	//	Subjects: []string{"math", "programming"},
	//}
	//fmt.Println("INSERT", st)
	//
	//err = students.Insert(st)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	// FIND
	fmt.Println("FIND")
	query := bson.M{}
	s := []student{}
	err = students.Find(query).All(&s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range s{
		fmt.Println(v)
	}

	// UPDATE
	fmt.Println("UPDATE")
	filter := bson.M{"age":44}
	change := bson.M{"$set":bson.M{"name":"old guy"}}
	s = []student{}
	_, err = students.UpdateAll(filter, change)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	query = bson.M{}
	s = []student{}
	err = students.Find(query).All(&s)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	for _, v := range s{
		fmt.Println(v)
	}
	// REMOVE
	fmt.Println("REMOVE")
	filter = bson.M{"name":"Alex"}
	_, err = students.RemoveAll(filter)

	query = bson.M{}
	s = []student{}
	err = students.Find(query).All(&s)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	for _, v := range s{
		fmt.Println(v)
	}
}