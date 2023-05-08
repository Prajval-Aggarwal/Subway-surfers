package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Struct for registration of the user
type RegisterRequest struct {
	P_Name   string `json:"playerName" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Struct for login request of the user
type LoginRequest struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

// Struct for update password request
type UpdatePasswordRequest struct {
	Password string `json:"password" `
}

// Struct for update name request
type UpdateNameRequest struct {
	P_Name string `json:"playerName" `
}

// Struct for forgot password request
type ForgotPassRequest struct {
	Email string `json:"email" `
}

//Validations on structs

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
