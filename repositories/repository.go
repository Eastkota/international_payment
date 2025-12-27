package repositories

import (
	"context"
	
	"payment_service/model"
)

type Repository interface {
    CreateMercRequest(ctx context.Context, input model.MercReqInput, transactionID string, purchaseDate string, mac string) error
}
