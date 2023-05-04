package model

type Payment struct {
	PaymentId   string `json:"primaryKey"`
	CartId      string
	PaymentType string
	Amount      float64
	Status      string
}
