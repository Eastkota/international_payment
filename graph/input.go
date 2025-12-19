package schema

import (
	"stripe_service/graph/scalar"

	"github.com/graphql-go/graphql"
)

var StripeInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "StripeInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"amount": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"transaction_id": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"user_id ": &graphql.InputObjectFieldConfig{
				Type: scalar.UUID,
			},
			"membership_duration_id ": &graphql.InputObjectFieldConfig{
				Type: scalar.UUID,
			},
		},
	},
)

var PaymentIntentInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "PaymentIntentInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"currency": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})

var StorePaymentIntentInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "StorePaymentIntentInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.Float,
		},
		"transaction_id": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"user_id": &graphql.InputObjectFieldConfig{
			Type: scalar.UUID,
		},
		"product": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"membership_duration_id": &graphql.InputObjectFieldConfig{
			Type: scalar.UUID,
		},
		"status": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
		"remarks": &graphql.InputObjectFieldConfig{
			Type: graphql.String,
		},
	},
})
