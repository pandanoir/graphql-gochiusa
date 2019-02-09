package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"./query"
	"github.com/graphql-go/graphql"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name:   "Query",
				Fields: graphql.Fields{"shop": &query.ShopField},
			},
		),
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	// GraphQLのshopフィールドの指定を行う
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Receive request\n")

	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	query := bufBody.String()

	result := executeQuery(query, schema)
	json.NewEncoder(w).Encode(result)
}

func main() {
	fmt.Printf("Start server\n")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
