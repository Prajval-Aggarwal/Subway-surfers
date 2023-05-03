package response

type CartResponse struct {
	CartId   string
	CartItem []CartItem
}

type CartItem struct {
	ItemId   string
	Quantity int64
}
