package repositories

import (
	"context"
	"time"

	"payment_service/model"
    "payment_service/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentRepository struct{
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

// repo.go
func (repo *PaymentRepository) CreateMercRequest(
    ctx context.Context, 
    input model.MercReqInput, 
    transactionID string, 
    purchaseDate string, 
    mac string,
) error {
    payment := &model.MercRequest{
        ID:                     uuid.New(),
        TransactionID:          transactionID,
        MerchantID:             config.MerchantId(),
        TransactionAmount:      input.TransactionAmount,
        TransactionCurrency:    input.TransactionCurrency,
        TransactionType:        "SALE",
        PurchaseDate:           time.Now(), // Or parse purchaseDate string
        MerchantTransactionMac: mac,
        Email:                  input.Email,
        CreatedAt:              time.Now(),
    }
    return repo.DB.WithContext(ctx).Create(payment).Error
}
