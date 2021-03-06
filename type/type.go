package gochiusa_type

import (
	"gopkg.in/mgo.v2/bson"
)

// ココは実際のドキュメントのフィールドを記述する。
// そのため、Shopはmembersフィールドは持たないが、memberListは持っている。
type Person struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}
type Shop struct {
	ID         bson.ObjectId   `bson:"_id"`
	Name       string          `bson:"name"`
	MemberList []bson.ObjectId `bson:"memberList"`
}
