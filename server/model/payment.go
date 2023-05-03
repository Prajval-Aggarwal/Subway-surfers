package model

type Payment struct {
	PaymentId   string
	CartId      string
	PaymentType string
	Amount      float64
	Status      string
}
