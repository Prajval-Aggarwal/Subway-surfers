package handler

import (
	"github.com/gin-gonic/gin"

	"subway/server/services/cart"
)

//	@Description	Show cart to the player
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//	@Tags			Cart
//	@Router			/show-cart [get]
func ShowCartHandler(ctx *gin.Context) {
	cart.ShowCartService(ctx)
}
