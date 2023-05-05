package request

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpdateAvatarRequest struct {
	AvatarId string
}

func (a UpdateAvatarRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AvatarId, validation.Required),
	)
}
