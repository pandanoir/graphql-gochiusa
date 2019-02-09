package main

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}
type Shop struct {
	Id         bson.ObjectId   `bson:"_id"`
	Name       string          `bson:"name"`
	MemberList []bson.ObjectId `bson:"memberList"`
}

func main() {
	session, _ := mgo.Dial("mongodb://localhost:27017")
	defer session.Close()
	db := session.DB("example")
	col := db.C("shop")

	col.RemoveAll(nil)
	chiya := Person{
		Id:   bson.NewObjectId(),
		Name: "宇治松千夜",
		Age:  16,
	}
	db.C("members").Insert(chiya)
	amausa := &Shop{
		Id:         bson.NewObjectId(),
		Name:       "甘兎庵",
		MemberList: []bson.ObjectId{chiya.Id},
	}

	sharo := Person{Id: bson.NewObjectId(), Name: "桐間紗路", Age: 16}
	db.C("members").Insert(sharo)
	fleur := &Shop{
		Id:         bson.NewObjectId(),
		Name:       "フルール・ド・ラパン",
		MemberList: []bson.ObjectId{sharo.Id},
	}

	cocoa := Person{Id: bson.NewObjectId(), Name: "保登心愛", Age: 16}
	chino := Person{Id: bson.NewObjectId(), Name: "香風智乃", Age: 14}
	rize := Person{Id: bson.NewObjectId(), Name: "天々座理世", Age: 17}
	db.C("members").Insert(cocoa)
	db.C("members").Insert(chino)
	db.C("members").Insert(rize)
	rabbithouse := &Shop{
		Id:         bson.NewObjectId(),
		Name:       "ラビットハウス",
		MemberList: []bson.ObjectId{cocoa.Id, chino.Id, rize.Id},
	}

	if err := col.Insert(fleur); err != nil {
		fmt.Printf("%+v \n", err)
	}
	if err := col.Insert(amausa); err != nil {
		fmt.Printf("%+v \n", err)
	}
	if err := col.Insert(rabbithouse); err != nil {
		log.Fatalln(err)
	}
}
