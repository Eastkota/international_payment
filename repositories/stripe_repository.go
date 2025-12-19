package repositories

import (
	"context"
	"fmt"
	"time"

	"stripe_service/model"

	"gorm.io/gorm"
)

type StripeRepository struct{
	DB *gorm.DB
}

func NewStripeRepository(db *gorm.DB) *StripeRepository {
	return &StripeRepository{DB: db}
}

func (repo *StripeRepository) StorePaymentIntent(ctx context.Context, inputData model.StorePaymentIntentInput) error {
    stripeResp := &model.StripeResponse{
        Amount:                fmt.Sprintf("%.2f", inputData.Amount),
        TransactionId:         inputData.TransactionId,
        UserId:                inputData.UserId,
        MembershipDurationId:  inputData.MembershipDurationId,
        Status:                inputData.Status,
        Remarks:               inputData.Remarks,
		Product:               inputData.Product,
        CreatedAt:             time.Now(),
        UpdatedAt:             time.Now(),
    }

    result := repo.DB.Create(stripeResp)
    if result.Error != nil {
        return fmt.Errorf("failed to save the stripe responses: %v", result.Error)
    }
    return nil
}
