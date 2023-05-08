package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Struct for power up
type PowerUpRequest struct {
	PowerUp_Id string `json:"powerUpId"`
}

// Struct for buying power up
type BuyRequest struct {
	PowerUp_Id string `json:"powerUpId" `
	Quantity   int64  `json:"quantity"`
}

// Validation of structs
func (a PowerUpRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.PowerUp_Id, validation.Required),
	)
}

func (a BuyRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.PowerUp_Id, validation.Required),
		validation.Field(&a.Quantity, validation.Required),
	)
}
