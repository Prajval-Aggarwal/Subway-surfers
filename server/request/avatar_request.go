package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Struct for update avatar request
type UpdateAvatarRequest struct {
	AvatarId string
}

// Validation of struct
func (a UpdateAvatarRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AvatarId, validation.Required),
	)
}
