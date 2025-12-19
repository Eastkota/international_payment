package services

import (
	"context"
	"stripe_service/model"
)

type Services interface {
	CreatePaymentIntent(ctx context.Context, inputData model.PaymentIntentInput) (*model.Stripe, error)
	StorePaymentIntent(ctx context.Context, inputData model.StorePaymentIntentInput) (error)
}
