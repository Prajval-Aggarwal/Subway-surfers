package handler

import (
	"subway/server/request"
	"subway/server/utils"

	"subway/server/services/payment"

	"github.com/gin-gonic/gin"
)

// @Description	Make payment
// @Accept			json
// @Produce		json
// @Success		200				{object}	response.Success
// @Failure		400				{object}	response.Error
// @Param			paymentDetails	body		request.PaymentRequest	true	"payment details of the player"
// @Tags			Authentication
// @Router			/make-payment [post]
func MakePaymentHandler(ctx *gin.Context) {
	// playerID, exists := ctx.Get("playerId")
	// fmt.Println("player id is :", playerID)
	// if !exists {
	// 	response.ErrorResponse(ctx, 401, "Unauthorised")
	// 	return
	// }
	var paymentRequest request.PaymentRequest
	utils.RequestDecoding(ctx, &paymentRequest)
	payment.MakePaymentService(ctx, "123", paymentRequest)

}
