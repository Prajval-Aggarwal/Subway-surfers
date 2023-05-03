package handler

import (
	"github.com/gin-gonic/gin"

	"subway/server/services/cart"
)

func ShowCartHandler(ctx *gin.Context) {
	cart.ShowCartService(ctx)
}
