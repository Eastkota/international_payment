package schema

import "github.com/graphql-go/graphql"

var GenericPaymentSuccessData = graphql.NewObject(graphql.ObjectConfig{
	Name: "GenericPaymentSuccessData",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
	},
})

var MercRequest = graphql.NewObject(graphql.ObjectConfig{
	Name: "MercRequest",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"transaction_type": &graphql.Field{Type: graphql.String},
		"merchant_id": &graphql.Field{Type: graphql.String},
		"transaction_amount": &graphql.Field{Type: graphql.String},
		"transaction_currency": &graphql.Field{Type: graphql.String},
		"transaction_id": &graphql.Field{Type: graphql.String},
		"original_transaction_id": &graphql.Field{Type: graphql.String},
		"purchase_date": &graphql.Field{Type: graphql.String},
		"email": &graphql.Field{Type: graphql.String},
		"merchant_transaction_mac": &graphql.Field{Type: graphql.String},
		"product": &graphql.Field{Type: graphql.String},
		"created_at": &graphql.Field{Type: graphql.String},
		"updated_at": &graphql.Field{Type: graphql.String},
	},
})

var KeyValueType = graphql.NewObject(graphql.ObjectConfig{
    Name: "KeyValuePair",
    Fields: graphql.Fields{
        "key":   &graphql.Field{Type: graphql.String},
        "value": &graphql.Field{Type: graphql.String},
    },
})

var InternationalPaymentResult = graphql.NewObject(graphql.ObjectConfig{
    Name: "InternationalPaymentResult",
    Fields: graphql.Fields{
        "acsURL": &graphql.Field{Type: graphql.String},
        // Since Payload is a map, we use a Scalar or a specific Object
        "payload": &graphql.Field{
            Type: graphql.NewList(KeyValueType), 
        },
		"error": &graphql.Field{
			Type: PaymentError,
		},
    },
})

