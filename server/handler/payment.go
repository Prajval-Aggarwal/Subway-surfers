package handler

import (
	"fmt"
	"subway/server/request"
	"subway/server/response"
	"subway/server/utils"

	"subway/server/services/payment"

	"github.com/gin-gonic/gin"
)

// @Description	Make payment
// @Accept			json
// @Produce		json
// @Success		200				{object}	response.Success
// @Failure		400				{object}	response.Error
// @Failure		401	{object}	response.Error
// @Param			paymentDetails	body		request.PaymentRequest	true	"payment details of the player"
// @Tags			Payment
// @Router			/make-payment [post]
func MakePaymentHandler(ctx *gin.Context) {
	playerID, exists := ctx.Get("playerId")
	fmt.Println("player id is :", playerID)
	if !exists {
		response.ErrorResponse(ctx, utils.UNAUTHORIZED, "Unauthorised")
		return
	}
	var paymentRequest request.PaymentRequest
	err := utils.RequestDecoding(ctx, &paymentRequest)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	err = paymentRequest.Validate()
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	payment.MakePaymentService(ctx, playerID.(string), paymentRequest)

}
