package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type PaymentRequest struct {
	CartId      string `json:"cartId"`
	PaymentType string `json:"paymentType"`
}

func (pr PaymentRequest) Validate() error {
	return validation.ValidateStruct(&pr,
		validation.Field(&pr.CartId, validation.Required),
		validation.Field(&pr.PaymentType, validation.Required),
	)
}
