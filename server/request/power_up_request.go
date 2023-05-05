package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PowerUpRequest struct {
	PowerUp_Id string `json:"powerUpId"`
}

type BuyRequest struct {
	PowerUp_Id string `json:"powerUpId" `
	Quantity   int64  `json:"quantity"`
}

func (pur PowerUpRequest) Validate() error {
	return validation.ValidateStruct(&pur,
		validation.Field(&pur.PowerUp_Id, validation.Required),
	)
}

func (br BuyRequest) Validate() error {
	return validation.ValidateStruct(&br,
		validation.Field(&br.PowerUp_Id, validation.Required),
		validation.Field(&br.Quantity, validation.Required),
	)
}
