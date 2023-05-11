package powerups

import (
	"subway/server/db"
	"subway/server/model"
	"subway/server/request"
	"subway/server/response"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

// ShowPowerUpService shows the list of all power ups
func ShowPowerUpsService(ctx *gin.Context) {
	query := "SELECT * FROM power_ups;"
	var powerups []model.PowerUp

	err := db.RawQuery(query, &powerups)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	response.ShowResponse("Sucess", utils.SUCCESS, "Power ups fetchedd sucessfully", powerups, ctx)

}

// UsePowerUpService uses the powerup and decreases the count for that powerup in players acccount
func UsePowerUpService(ctx *gin.Context, playerID string, powerupRequest request.PowerUpRequest) {
	query := "SELECT * FROM player_power_ups WHERE p_id=? AND power_up_id=?;"
	var playerPowerup model.PlayerPowerUps
	if powerupRequest.PowerUp_Id == "" {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Please provide valid id")
		return
	}
	err := db.RawQuery(query, &playerPowerup, playerID, powerupRequest.PowerUp_Id)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	//Chckingg whether the player has that powerup or not
	if playerPowerup.Quantity == 0 {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Player do not have this powerup")
		return
	}

	//Decreasing powerup count
	playerPowerup.Quantity = playerPowerup.Quantity - 1

	query = "UPDATE player_power_ups SET quantity=? WHERE p_id=? AND power_up_id=?;"

	err = db.ExecuteQuery(query, playerPowerup.Quantity, playerID, powerupRequest.PowerUp_Id)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	response.ShowResponse("Success", utils.SUCCESS, "Updated sucess fully", nil, ctx)
}

// BuyPowerupService adds the powerups to the player account
func BuyPowerupService(ctx *gin.Context, playerId string, BuyRequest request.BuyRequest) {

	//find the power up with that id calculate total amount
	var powerUp model.PowerUp
	var playerCoins model.PlayerCoins
	var playerPowerups model.PlayerPowerUps
	var exists bool

	if BuyRequest.PowerUp_Id == "" {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Please provide valid id")
		return
	}

	err := db.FindById(&powerUp, BuyRequest.PowerUp_Id, "power_up_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	amount := powerUp.Price * BuyRequest.Quantity

	//decrease the coins of the player
	err = db.FindById(&playerCoins, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	//Check if the player has enough coins to buy taht powerup or not
	if playerCoins.Coins-amount < 0 {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Not enough coin to buy powerup")
		return
	}

	playerCoins.Coins = playerCoins.Coins - amount
	err = db.UpdateRecord(&playerCoins, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	//add the power ups to player_powerups table add or update

	query := "SELECT EXISTS(SELECT * FROM player_power_ups WHERE p_id=? AND power_up_id=?)"
	err = db.RawQuery(query, &exists, playerId, BuyRequest.PowerUp_Id)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	if !exists {
		playerPowerups.PowerUp_Id = BuyRequest.PowerUp_Id
		playerPowerups.Quantity = BuyRequest.Quantity
		db.CreateRecord(&playerPowerups)
	} else {

		query = "SELECT * FROM player_power_ups WHERE p_id=? AND power_up_id=?;"
		db.RawQuery(query, &playerPowerups, playerId, BuyRequest.PowerUp_Id)

		playerPowerups.Quantity = playerPowerups.Quantity + BuyRequest.Quantity

		query = "UPDATE player_power_ups SET quantity=? WHERE p_id=? AND power_up_id=?;"

		err = db.ExecuteQuery(query, playerPowerups.Quantity, playerId, BuyRequest.PowerUp_Id)
		if err != nil {
			response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
			return
		}

	}

}
