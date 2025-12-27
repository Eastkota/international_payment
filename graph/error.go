package schema

import "github.com/graphql-go/graphql"

var PaymentError = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentError",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
		"code":    &graphql.Field{Type: graphql.String},
		"field":   &graphql.Field{Type: graphql.String},
	},
})
