package types

import (
	"github.com/graphql-go/graphql"
)

//NotToDo represents something we're not going to do.
var NotTodo = graphql.NewObject(graphql.ObjectConfig {
	Name: "NotTodo",
	Fields: graphql.Fields {
		"name": &graphql.Field {
			Type: graphql.String,
		},
		"description": &graphql.Field {
			Type: graphql.String,
		},
	},
})