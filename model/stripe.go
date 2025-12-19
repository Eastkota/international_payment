package model

import (
	"time"
	"github.com/google/uuid"
)

type Customer struct {
	Email string `gorm:"type:varchar(100)" json:"email"`
}

type Stripe struct {
	TransactionId string `gorm:"type:text" json:"transaction_id"`
	ClientSecret  string `gorm:"type:text" json:"client_secret"`
	PublishedKey  string `gorm:"type:text" json:"published_key"`
}
type StripeResult struct {
	Stripe *Stripe `json:"stripe"`
}
type GenericStripeSuccessData struct {
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

type StripeResponse struct {
	Amount                string    `gorm:"type:varchar" json:"amount"`
	TransactionId         string    `gorm:"type:varchar" json:"transaction_id"`
	UserId                uuid.UUID `gorm:"type:uuid" json:"user_id"`
	MembershipDurationId  uuid.UUID `gorm:"type:uuid" json:"membership_duration_id"`
	Status                string    `gorm:"type:varchar" json:"status"`
	Remarks               string    `gorm:"type:varchar" json:"remarks"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Product               string    `gorm:"type:varchar" json:"product"`
}

func (StripeResponse) TableName() string {
    return "payment.payment_responses_stripe"
}
