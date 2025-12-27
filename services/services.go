package services

import (
	"context"
	"payment_service/model"
)

type Services interface {
	CreateInternationalPayment(ctx context.Context, inputData model.MercReqInput) (*model.InternationalPaymentResult, error)
}
