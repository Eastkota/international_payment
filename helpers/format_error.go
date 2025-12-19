package helpers

import "stripe_service/model"

func FormatError(err error) *model.GenericStripeResponse {
	return &model.GenericStripeResponse{
		Data: nil,
		Error: &model.StripeError{
			Message: err.Error(),
		},
	}
}
