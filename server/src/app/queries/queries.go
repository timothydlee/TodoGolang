package queries

import (
	"github.com/graphql-go/graphql"
	fields "app/queries/fields"
)

// RootQuery is the root query.
var RootQuery = graphql.NewObject(graphql.ObjectConfig {
	Name: "RootQuery",
	Fields: graphql.Fields {
		"getNotTodos": fields.GetNotTodos,
	},
})