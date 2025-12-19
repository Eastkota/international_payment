package model

import (
	"github.com/google/uuid"
)

type StripeInput struct {
	MembershipDurationId uuid.UUID `json:"membership_duration_id"`
	Email                string `json:"email"`
	CardNo               string `json:"card_no"`
	ExperationDate       string `json:"experation_date"`
	SecurityCode         string `json:"security_code"`
	Country              string `json:"country"`
}

type PaymentIntentInput struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
type StorePaymentIntentInput struct {
	Amount               float64 `json:"amount"`
	TransactionId        string  `json:"transaction_id"`
	UserId               uuid.UUID  `json:"user_id"`
	MembershipDurationId uuid.UUID  `json:"membership_duration_id"`
	Status               string  `json:"status"`
	Remarks              string  `json:"remarks"`
	Product              string  `json:"product"`
}
