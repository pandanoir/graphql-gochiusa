package query

import (
	"fmt"

	gochiusa_type "../type"
	"github.com/graphql-go/graphql"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var stringField = &graphql.Field{Type: graphql.String}
var MemberType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Member",
		Fields: graphql.Fields{
			"id":   stringField,
			"name": stringField,
			"age":  &graphql.Field{Type: graphql.Int},
		},
	},
)
var shopType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Shop",
		Fields: graphql.Fields{
			"id":   stringField,
			"name": stringField,
			"members": &graphql.Field{
				Type: graphql.NewList(MemberType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					shop, _ := p.Source.(*(gochiusa_type.Shop))

					session, _ := mgo.Dial("mongodb://localhost:27017")
					defer session.Close()
					db := session.DB("example")
					DB := db.C("members")

					// nameがnameQueryになっているレコードを探す
					members := []gochiusa_type.Person{}
					for _, memberidQuery := range shop.MemberList {
						member := gochiusa_type.Person{}
						if err := DB.Find(bson.M{"_id": memberidQuery}).One(&member); err != nil {
							fmt.Printf("%+v \n", err)
						}
						members = append(members, member)
					}
					return &members, nil
				},
			},
		},
	},
)

// define query schema
var ShopField = graphql.Field{
	Type: shopType,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		nameQuery, isOK := p.Args["name"].(string)
		if !isOK {
			return nil, nil
		}

		session, _ := mgo.Dial("mongodb://localhost:27017")
		defer session.Close()
		db := session.DB("example")
		DB := db.C("shop")

		// nameがnameQueryになっているレコードを探す
		s := gochiusa_type.Shop{}
		if err := DB.Find(bson.M{"name": nameQuery}).One(&s); err != nil {
			fmt.Printf("%+v \n", err)
		}

		return &s, nil
	},
}
