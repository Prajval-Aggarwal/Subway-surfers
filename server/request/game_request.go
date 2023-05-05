package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type EndGameRequest struct {
	CoinsCollected int64 `json:"coinsCollected" `
	Distance       int64 `json:"distance" `
}

func (egr EndGameRequest) Validate() error {
	return validation.ValidateStruct(&egr,
		validation.Field(&egr.CoinsCollected, validation.Required),
		validation.Field(&egr.Distance, validation.Required),
	)
}
