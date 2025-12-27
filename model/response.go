package model

type PaymentSuccessData struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type GenericPaymentResponse struct {
	Data  interface{}
	Error *PaymentError
}

type MkKeyRequest struct {
    MerchantID string `json:"merchantId"`
    PubKey     string `json:"pubKey"`
    PurchaseID string `json:"purchaseId"`
}

type MkKeyResponse struct {
    MerchantID string `json:"merchantId"`
    PurchaseID string `json:"purchaseId"`
    PubKey     string `json:"pubKey"`
    ErrorCode  string `json:"errorCode"`
}


