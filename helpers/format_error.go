package helpers

import "payment_service/model"

func FormatError(err error) *model.GenericPaymentResponse {
	return &model.GenericPaymentResponse{
		Data: nil,
		Error: &model.PaymentError{
			Message: err.Error(),
		},
	}
}

func FormatInternationalPaymentError(err error) *model.GenericPaymentResponse {
	return &model.GenericPaymentResponse{
		Data: nil,
		Error: &model.PaymentError{
			Message: err.Error(),
		},
	}
}
