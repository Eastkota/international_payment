package repositories

import (
	"context"
	"stripe_service/model"
)

type Repository interface {
	StorePaymentIntent(ctx context.Context, inputData model.StorePaymentIntentInput) (error)
}
