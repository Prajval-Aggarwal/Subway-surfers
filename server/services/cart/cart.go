package cart

import (
	"fmt"
	"math/rand"
	"subway/server/db"
	"subway/server/model"
	"subway/server/response"
	"subway/server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// GenerateCart generates cart for the player
func GenerateCart() {
	var cart model.Cart
	now := time.Now()
	cart.CreatedAt = now
	err := db.CreateRecord(&cart)
	if err != nil {
		return
	}

	// Selecting a function randomly
	fn := SelectRandomFunc()
	if fn != nil {
		fn(cart.CartId)
	}
}

func SelectRandomFunc() func(string) {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(2)
	switch randNum {
	case 0:
		return AddPowerupsToCart
	case 1:
		return AddCoinsToCart
	default:
		return nil
	}
}

// Choosing random power ups to add to cart
func AddPowerupsToCart(cartId string) {
	fmt.Println("cart id is", cartId)
	fmt.Println("add to power ups function called")
	var powerUps []struct {
		PowerUpID string
	}
	query := "SELECT power_up_id FROM power_ups ORDER BY RANDOM()  LIMIT 3;"
	db.RawQuery(query, &powerUps)
	for _, powerup := range powerUps {
		var item model.CartItem
		item.CartId = cartId
		item.ItemId = powerup.PowerUpID
		item.Quantity = int64(rand.Intn(20-10) + 10)
		db.CreateRecord(&item)
	}

}

// Choosing random amount of coins to add to cart
func AddCoinsToCart(cartId string) {
	item := model.CartItem{
		CartId:   cartId,
		ItemId:   "6",
		Quantity: int64(rand.Intn(100000-10000) + 10000),
	}
	db.CreateRecord(&item)
}

// ShowCartService shows the cart generated for that day
func ShowCartService(ctx *gin.Context) {
	var cartDetails response.CartResponse
	var items []response.CartItem
	query := "SELECT cart_id FROM carts ORDER BY created_at DESC LIMIT 1"
	cartDetails = db.Fun1(query)

	query1 := "SELECT item_id,quantity FROM cart_items WHERE cart_id =?"
	err := db.RawQuery(query1, &items, cartDetails.CartId)
	if err != nil {
		response.ErrorResponse(ctx, utils.INTERNAL_SERVER_ERROR, err.Error())
		return
	}
	cartDetails.CartItem = items

	response.ShowResponse("Success", utils.SUCCESS, "cart data fetched successfully", cartDetails, ctx)
}
