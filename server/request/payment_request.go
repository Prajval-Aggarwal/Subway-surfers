package request

type PaymentRequest struct {
	CartId      string `json:"cartId"`
	PaymentType string `json:"paymentType"`
}
