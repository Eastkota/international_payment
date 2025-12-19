package resolvers

import (
	"stripe_service/helpers"
	"stripe_service/model"
	"stripe_service/services"

	"encoding/json"
	
	"github.com/graphql-go/graphql"
)

type StripeResolver struct {
	Services *services.StripeService
}

func NewStripeResolver(service *services.StripeService) *StripeResolver {
	return &StripeResolver{Services: service}
}

func (r *StripeResolver) CreatePaymentIntent(p graphql.ResolveParams) *model.GenericStripeResponse {
	var stripeInput model.PaymentIntentInput
	inputData := p.Args["input"].(map[string]interface{})
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return helpers.FormatError(err)
	}

	err = json.Unmarshal(jsonData, &stripeInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	result, err := r.Services.CreatePaymentIntent(p.Context, stripeInput)
	if err != nil {
		return helpers.FormatError(err)
	}

	return &model.GenericStripeResponse{
		Data: &model.StripeResult{
			Stripe: result,
		},
		Error: nil,
	}
}

func (r *StripeResolver) StorePaymentIntent(p graphql.ResolveParams) *model.GenericStripeResponse {

	var storeStripeInput model.StorePaymentIntentInput

	inputData := p.Args["input"].(map[string]interface{})
	jsonData, err := json.Marshal(inputData)
	if err != nil {
		return helpers.FormatError(err)
	}

	err = json.Unmarshal(jsonData, &storeStripeInput)
	if err != nil {
		return helpers.FormatError(err)
	}
	err = r.Services.StorePaymentIntent(p.Context, storeStripeInput)
	if err != nil {
		return helpers.FormatError(err)
	}
	return &model.GenericStripeResponse{
		Data: &model.GenericStripeSuccessData{
			Message: "Payment intent stored successfully",
		},
		Error: nil,
	}
}
