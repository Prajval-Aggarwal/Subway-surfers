package request

type RegisterRequest struct {
	P_Name   string `json:"playerName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=16,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LogoutRequest struct {
	P_Id string `json:"playerId" validate:"required"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" validate:"required"`
}

type UpdateNameRequest struct {
	P_Name string `json:"playerName" validate:"required"`
}

type ForgotPassRequest struct {
	Email string `json:"email" validate:"required,email"`
}
