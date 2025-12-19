package services

import (
	"context"
	"encoding/json"
	
	"stripe_service/config"
	"stripe_service/helpers"
	"stripe_service/model"
	"stripe_service/repositories"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
)

type StripeService struct{
	Repository repositories.Repository
}

func NewStripeService(repository repositories.Repository) *StripeService {
	return &StripeService{Repository: repository}
}

func (ps *StripeService) CreatePaymentIntent(ctx context.Context, inputData model.PaymentIntentInput) (*model.Stripe, error) {
	stripe.Key = config.ClientSecretKey()
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(helpers.DollarsToCents(inputData.Amount)),
		Currency:           stripe.String(string(inputData.Currency)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(false),
		},
	}
	result, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	// Unmarshal the JSON into a map
	var resultMap map[string]interface{}
	err = json.Unmarshal(data, &resultMap)
	if err != nil {
		return nil, err
	}

	return &model.Stripe{
		TransactionId: resultMap["id"].(string),
		ClientSecret:  resultMap["client_secret"].(string),
		PublishedKey: config.PublishedKey(),
	}, nil
}


func (ps *StripeService) StorePaymentIntent(ctx context.Context, inputData model.StorePaymentIntentInput) (error) {
	return ps.Repository.StorePaymentIntent(ctx, inputData)
}


