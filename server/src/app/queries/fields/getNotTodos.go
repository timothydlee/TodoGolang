package queries

import (
	"context"
	"github.com/graphql-go/graphql"
	bson "github.com/mongodb/mongo-go-driver/bson"

	"app/data"
	types "app/types"
)

type todoStruct struct {
	NAME		string `json:"name"`
	DESCRIPTION	string `json:'description"`
}

// GetNotTodos is the graphql field which we will pull notTodos from.
var GetNotTodos = &graphql.Field {
	Type:			graphql.NewList(types.NotTodo),
	Description:	"Get all the not todos",
	Resolve:		func(params graphql.ResolveParams) (interface{}, error) {
		notTodoCollection := mongo.Client.Database("medium-app").Collection("Not_Todos")
		todos, err := notTodoCollection.Find(context.Background(), nil)
		if err != nil { panic(err) }
		var todosList []todoStruct
		for todos.Next(context.Background()) {
			doc := bson.NewDocument()

			err := todos.Decode(doc)
			if err != nil { panic(err) }

			keys, err := doc.Keys(false)
			if err != nil { panic(err) }

			// Convert BSON to struct
			todo := todoStruct{}
			for _, key := range keys {
				keyString := key.String()
				elm, err := doc.Lookup(keyString)
				if err != nil { panic(err) }

				switch (keyString) {
				case "name":
					todo.NAME = elm.Value().StringValue()
				case "description":
					todo.DESCRIPTION = elm.Value().StringValue()
				default:
				}
			}
			todosList = append(todosList, todo)
		}
		
		return todosList, nil
	},
}