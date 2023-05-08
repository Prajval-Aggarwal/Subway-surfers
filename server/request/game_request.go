package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Struct for ending the game
type EndGameRequest struct {
	CoinsCollected int64 `json:"coinsCollected" `
	Distance       int64 `json:"distance" `
}

// Validation for struct
func (a EndGameRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.CoinsCollected, validation.Required),
		validation.Field(&a.Distance, validation.Required),
	)
}
