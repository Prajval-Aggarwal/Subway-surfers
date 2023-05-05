package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegisterRequest struct {
	P_Name   string `json:"playerName" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type UpdatePasswordRequest struct {
	Password string `json:"password" `
}

type UpdateNameRequest struct {
	P_Name string `json:"playerName" `
}

type ForgotPassRequest struct {
	Email string `json:"email" `
}

//validations added

func (a RegisterRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.P_Name, validation.Required, validation.Length(5, 15)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required),
	)
}

func (a LoginRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required),
	)
}

func (a UpdatePasswordRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Password, validation.Required),
	)
}

func (a UpdateNameRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.P_Name, validation.Required, validation.Length(5, 15)),
	)
}

func (a ForgotPassRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}
