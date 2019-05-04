package types

import (
	"github.com/graphql-go/graphql"
)

// This is NotTodo comment
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