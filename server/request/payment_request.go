package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Struct for Make payment
type PaymentRequest struct {
	CartId      string `json:"cartId"`
	PaymentType string `json:"paymentType"`
}

// Validation for Make payment
func (a PaymentRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.CartId, validation.Required),
		validation.Field(&a.PaymentType, validation.Required),
	)
}
