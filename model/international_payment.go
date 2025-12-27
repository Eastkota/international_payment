package model

import (
	"time"
	"github.com/google/uuid"
)

type GenericPaymentSuccessData struct {
	Message string `json:"string"`
}

type AuthUserMembership struct {
    ID                   uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    MembershipJoinDate   time.Time `gorm:"type:timestamptz;not null" json:"membership_join_date"`
    MembershipEndDate    time.Time `gorm:"type:timestamptz;not null" json:"membership_end_date"`
    UserId               uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    MembershipDurationId uuid.UUID `gorm:"type:uuid;not null" json:"membership_duration_id"`
    CreatedAt            time.Time `json:"created_at"`
    UpdatedAt            time.Time `json:"updated_at"`
}

type MPIKeyRequest struct {
    MerchantID string `json:"merchantId"` // Mandatory
    PurchaseID string `json:"purchaseId"` // Mandatory (Min 6 digits)
    PubKey     string `json:"pubKey"`     // Mandatory (RSA 2048, Base64Url)
}


type MercRequest struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
    TransactionType string `gorm:"type:varchar(50)" json:"transaction_type"`
    MerchantID      string `gorm:"type:varchar(100)" json:"merchant_id"`
    TransactionAmount string `gorm:"type:varchar(50)" json:"transaction_amount"`
    TransactionCurrency string `gorm:"type:varchar(10)" json:"transaction_currency"`
    TransactionID   string `gorm:"type:varchar(100)" json:"transaction_id"`
    OriginalTransactionID string `gorm:"type:varchar(100)" json:"original_transaction_id"`
    PurchaseDate    time.Time `gorm:"type:timestamptz" json:"purchase_date"`
    Email          string `gorm:"type:varchar(100)" json:"email"`
    MerchantTransactionMac string `gorm:"type:varchar(255)" json:"merchant_transaction_mac"`
    Product       string `gorm:"type:varchar(255)" json:"product"`
    CardHolderName string `gorm:"type:varchar(100)" json:"card_holder_name"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

func (MercRequest) TableName() string {
    return "payment.international_payment"
}

type InternationalPaymentResult struct {
    AcsURL  string            `json:"acs_url"`
    Payload map[string]string `json:"payload"`
    Error   *PaymentError     `json:"error,omitempty"`
}

