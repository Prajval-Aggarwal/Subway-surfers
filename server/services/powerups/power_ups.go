package powerups

import (
	"subway/server/db"
	"subway/server/model"
	"subway/server/request"
	"subway/server/response"

	"github.com/gin-gonic/gin"
)

func ShowPowerUpsService(ctx *gin.Context) {
	query := "SELECT * FROM power_ups;"
	var powerups []model.PowerUp

	err := db.RawQuery(query, powerups)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	response.ShowResponse("Sucess", 200, "Power ups fetchedd sucessfully", powerups, ctx)

}

func UsePowerUpService(ctx *gin.Context, playerID string, powerupRequest request.PowerUpRequest) {
	query := "SELECT * FROM player_power_ups WHERE p_id=? AND power_up_id=?;"
	var playerPowerup model.PlayerPowerUps
	err := db.RawQuery(query, playerPowerup, playerID, powerupRequest.PowerUp_Id)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	if playerPowerup.Quantity == 0 {
		response.ErrorResponse(ctx, 400, "Player do not have this powerup")
		return
	}

	playerPowerup.Quantity = playerPowerup.Quantity - 1

	query = "UPDATE player_power_ups SET quantity=? WHERE p_id=? AND power_up_id=?;"

	err = db.ExecuteQuery(query, playerPowerup.Quantity, playerID, powerupRequest.PowerUp_Id)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	response.ShowResponse("Success", 200, "Updated sucess fully", nil, ctx)
}

func BuyPowerupService(ctx *gin.Context, playerId string, BuyRequest request.BuyRequest) {

	//find the power up with that id calculate total amount
	var powerUp model.PowerUp
	var playerCoins model.PlayerCoins
	err := db.FindById(&powerUp, BuyRequest.PowerUp_Id, "power_up_id")
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	amount := powerUp.Price * BuyRequest.Quantity

	//decrease the coins o the player and update the record
	db.FindById(&playerCoins, playerId, "p_id")

	if playerCoins.Coins-amount < 0 {
		response.ErrorResponse(ctx, 400, "Not enough coin to buy powerup")
		return
	}

	playerCoins.Coins = playerCoins.Coins - amount
	err = db.UpdateRecord(&playerCoins, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	//add the power ups to player_powerups table add or update
	var playerPowerups model.PlayerPowerUps
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM player_power_ups WHERE p_id=? AND power_up_id=?)"
	db.RawQuery(query, &exists, playerId, BuyRequest.PowerUp_Id)
	if !exists {
		playerPowerups.PowerUp_Id = BuyRequest.PowerUp_Id
		playerPowerups.Quantity = BuyRequest.Quantity
		db.CreateRecord(&playerPowerups)
	} else {
		playerPowerups.Quantity += BuyRequest.Quantity
		query = "UPDATE player_power_ups SET quantity=? WHERE p_id=? AND power_up_id=?;"

		err = db.ExecuteQuery(query, playerPowerups.Quantity, playerId, BuyRequest.PowerUp_Id)
		if err != nil {
			response.ErrorResponse(ctx, 400, err.Error())
			return
		}

	}

}
