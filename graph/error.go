package schema

import "github.com/graphql-go/graphql"

var StripeError = graphql.NewObject(graphql.ObjectConfig{
	Name: "StripeError",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"code":    &graphql.Field{Type: graphql.String},
		"field":   &graphql.Field{Type: graphql.String},
	},
})
