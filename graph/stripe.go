package schema

import "github.com/graphql-go/graphql"

var Stripe = graphql.NewObject(graphql.ObjectConfig{
	Name: "Stripe",
	Fields: graphql.Fields{
		"transaction_id": &graphql.Field{Type: graphql.String},
		"client_secret":  &graphql.Field{Type: graphql.String},
		"published_key":  &graphql.Field{Type: graphql.String},
	},
})
var StripeResult = graphql.NewObject(graphql.ObjectConfig{
	Name: "StripeResult",
	Fields: graphql.Fields{
		"stripe": &graphql.Field{Type: Stripe},
	},
})

var GenericStripeSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericStripeSuccessData",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
	},
})

// var GenericStripeSuccessDataResult = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "GenericStripeSuccessDataResult",
// 	Fields: graphql.Fields{
// 		"generic_stripe_success_data": &graphql.Field{Type: GenericStripeSuccessData},
// 	},
// })