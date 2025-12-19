package model

type StripeSuccessData struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type GenericStripeResponse struct {
	Data  interface{}
	Error *StripeError
}
