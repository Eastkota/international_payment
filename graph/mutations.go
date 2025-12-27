package schema

import (
	"payment_service/resolvers"

	"github.com/graphql-go/graphql"
)

func NewMutationType(resolver *resolvers.PaymentResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createInternationalPayment": &graphql.Field{
				Type:        InternationalPaymentResult,
				Description: "Create International Payment",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: MercRequestInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Extract input from arguments
					return resolver.CreateInternationalPayment(p)
				},
			},
		},
	})
}		
