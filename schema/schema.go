package gochiusa

import (
	"../query"
	"github.com/graphql-go/graphql"
)

func InitSchema() graphql.Schema {
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"shop": &query.ShopField,
			},
		},
	)
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	return schema
}
