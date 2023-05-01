package request

type EndGameRequest struct {
	CoinsCollected int64 `json:"coinsCollected" validate:"required"`
	Distance       int64 `json:"distance" validate:"required"`
}
