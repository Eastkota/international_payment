package schema

import "github.com/graphql-go/graphql"

var StripeResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "StripeResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: StripeResult,
		},
		"error": &graphql.Field{
			Type: StripeError,
		},
	},
})

var GenericStripeSuccessResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericStripeSuccessResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: GenericStripeSuccessData,
		},
		"error": &graphql.Field{
			Type: StripeError,
		},
	},
})
