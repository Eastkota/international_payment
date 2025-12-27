package schema

import "github.com/graphql-go/graphql"

var PaymentResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "PaymentResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: InternationalPaymentResult,
		},
		"error": &graphql.Field{
			Type: PaymentError,
		},
	},
})

var GenericPaymentSuccessResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericPaymentSuccessResponse",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type: InternationalPaymentResult, // Changed from GenericPaymentSuccessData
		},
		"error": &graphql.Field{
			Type: PaymentError,
		},
	},
})
