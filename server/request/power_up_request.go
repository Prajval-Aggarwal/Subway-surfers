package request

type PowerUpRequest struct {
	PowerUp_Id string `json:"powerUpId" validate:"required"`
}

type BuyRequest struct {
	PowerUp_Id string `json:"powerUpId" validate:"required"`
	Quantity   int64  `json:"quantity" validate:"required"`
}
