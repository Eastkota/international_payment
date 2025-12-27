package schema

import (
	"github.com/graphql-go/graphql"
)

var MercRequestInput = graphql.NewInputObject(graphql.InputObjectConfig{
    Name: "MercRequestInput",
    // Change this line right here:
    Fields: graphql.InputObjectConfigFieldMap{
        "transactionAmount": &graphql.InputObjectFieldConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
        "transactionCurrency": &graphql.InputObjectFieldConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
        "email": &graphql.InputObjectFieldConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
        "product": &graphql.InputObjectFieldConfig{
            Type: graphql.String,
        },
        "cardHolderName": &graphql.InputObjectFieldConfig{
            Type: graphql.String,
        },
    },
})

