package model

// What the Frontend sends you
type MercReqInput struct {
    TransactionAmount   string `json:"transactionAmount"`   // e.g. "100.00"
    TransactionCurrency string `json:"transactionCurrency"` // e.g. "840"
    Email               string `json:"email"`               // Student email
    Product             string `json:"product"`             // Course Name
    CardHolderName      string `json:"cardHolderName"`      // Student Name
}
