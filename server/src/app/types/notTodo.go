package types

import (
	"github.com/graphql-go/graphql"
)

// NotTodo mdoels our not-todo object. It has the name and description of 
// what we're not going to do
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