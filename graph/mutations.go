package schema

import (
	"stripe_service/resolvers"

	"github.com/graphql-go/graphql"
)

func NewMutationType(resolver *resolvers.StripeResolver) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createPaymentIntent": &graphql.Field{
				Type: StripeResponse,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: PaymentIntentInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Extract input from arguments
					return PublicAuthMiddleware(resolver.CreatePaymentIntent)(p), nil
				},
			},
			"storePaymentIntent": &graphql.Field{
				Type: GenericStripeSuccessResponse,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: StorePaymentIntentInput,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return PublicAuthMiddleware(resolver.StorePaymentIntent)(p), nil
				},
			},
		},
	})
}		
