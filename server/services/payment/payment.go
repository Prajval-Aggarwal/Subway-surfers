package payment

import (
	"fmt"
	"subway/server/db"
	"subway/server/model"
	"subway/server/request"
	"subway/server/response"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"

	"github.com/stripe/stripe-go/v72/paymentintent"
)

func StripePayment(amount float64, ctx *gin.Context) (pi, pi1 *stripe.PaymentIntent) {
	// stripe payment integration
	stripe.Key = "sk_test_51N3YlcSE1Cg6ZXrAgAvQRmxmKYiSijrryYAFpkZXVTSGebnkimNxgPVlYoQmy0EI9DwyKyEThIsxQZZHTHSSjNwg00zFyDkubC"

	// Get the amount from the request
	// amount := billamount
	fmt.Println("amount", amount)
	// Create a new PaymentIntent
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(int64(amount * 100)),
		Currency:           stripe.String("inr"),
		Description:        stripe.String("Payment"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	params1 := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}

	pi1, err = paymentintent.Confirm(pi.ID, params1)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	// Check the payment intent status
	switch pi1.Status {
	case "succeeded":
		// Payment succeeded
		response.ShowResponse("Success", 200, "Payment processed Successfully", "", ctx)
		return
	case "requires_payment_method":
		response.ErrorResponse(ctx, 400, "Requires Payment Method")
		return
	case "requires_action":
		// Additional action required
		if pi1.Status == "requires_action" && pi1.NextAction != nil {
			switch pi1.NextAction.Type {
			case "use_stripe_sdk":

				response.ShowResponse("Success", 200, "Payment processed Successfully , Here is your client secret", pi1, ctx)
			}
		}
	default:
		response.ErrorResponse(ctx, 400, "Payment requires more actions")
		return
	}

	return pi, pi1

}

func MakePaymentService(ctx *gin.Context, playerId string, paymentRequest request.PaymentRequest) {
	var paymentDetails model.Payment
	var cartItems []model.CartItem
	db.FindById(&cartItems, paymentRequest.CartId, "cart_id")
	paymentDetails.CartId = paymentRequest.CartId
	paymentDetails.PaymentType = paymentRequest.PaymentType

	totalAmount := float64(0)
	//calculate amount from the items present in the cart
	for _, x := range cartItems {
		if x.ItemId == "1" {
			totalAmount += float64(5 * x.Quantity)
		} else if x.ItemId == "2" {
			totalAmount += float64(10 * x.Quantity)
		} else if x.ItemId == "3" {
			totalAmount += float64(15 * x.Quantity)
		} else if x.ItemId == "4" {
			totalAmount += float64(10 * x.Quantity)
		} else if x.ItemId == "5" {
			totalAmount += float64(15 * x.Quantity)
		} else if x.ItemId == "6" {
			totalAmount += float64(0.0133 * float64(x.Quantity))
		} else {
			response.ErrorResponse(ctx, 400, "Error in cart items")
			return
		}
	}

	if totalAmount > 500 || totalAmount > 1000 {
		totalAmount = float64(totalAmount) - float64(totalAmount)*0.1
	}
	paymentDetails.Amount = totalAmount
	pi, pi1 := StripePayment(totalAmount, ctx)
	paymentDetails.PaymentId = pi.ID
	paymentDetails.Status = string(pi1.Status)
	//create payment record
	err := db.CreateRecord(&paymentDetails)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	//create player payment record
	playerPayment := model.PlayerPayment{
		P_ID:      playerId,
		PaymentId: pi.ID,
	}
	db.CreateRecord(&playerPayment)
	response.ShowResponse("Success", 200, string(pi.Status), paymentDetails, ctx)
}
