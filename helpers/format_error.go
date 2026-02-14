package helpers

import (
	"errors"
	"log"
	"payment_service/model"
)

func FormatError(err error) *model.GenericPaymentResponse {
	var appErr *AppError
	if errors.As(err, &appErr) {
		if appErr.Internal != nil {
			log.Printf("[ERROR] code=%s message=%q internal=%q", appErr.Code, appErr.Message, appErr.Internal.Error())
		}
		return &model.GenericPaymentResponse{
			Data: nil,
			Error: &model.PaymentError{
				Message: appErr.Message,
				Code:    string(appErr.Code),
				Field:   appErr.Field,
			},
		}
	}
	log.Printf("[ERROR] code=INTERNAL_ERROR untyped_error=%q", err.Error())
	return &model.GenericPaymentResponse{
		Data: nil,
		Error: &model.PaymentError{
			Message: err.Error(),
			Code:    string(ErrCodeInternal),
		},
	}
}

func FormatInternationalPaymentError(err error) *model.GenericPaymentResponse {
	var appErr *AppError
	if errors.As(err, &appErr) {
		if appErr.Internal != nil {
			log.Printf("[ERROR] code=%s message=%q internal=%q", appErr.Code, appErr.Message, appErr.Internal.Error())
		}
		return &model.GenericPaymentResponse{
			Data: nil,
			Error: &model.PaymentError{
				Message: appErr.Message,
				Code:    string(appErr.Code),
				Field:   appErr.Field,
			},
		}
	}
	log.Printf("[ERROR] code=INTERNAL_ERROR untyped_error=%q", err.Error())
	return &model.GenericPaymentResponse{
		Data: nil,
		Error: &model.PaymentError{
			Message: err.Error(),
			Code:    string(ErrCodeInternal),
		},
	}
}
