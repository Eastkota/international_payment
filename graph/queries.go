package schema

import (
	"payment_service/helpers"
	"payment_service/model"
	"payment_service/resolvers"

	"github.com/graphql-go/graphql"
)

func NewQueryType(resolver *resolvers.PaymentResolver) *graphql.Object { 
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"service": &graphql.Field{
				Type: graphql.NewNonNull(Service),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					schema, err := GetSchema()
					if err != nil {
						return nil, err
					}
					serviceInfo := model.Service{
						Name:    "StripeService",
						Version: "1.0.0",
						Schema:  helpers.ConvertSchemaToString(schema),
					}
					return serviceInfo, nil
				},
			},
		},
	})
}