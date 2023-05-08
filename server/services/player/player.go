package player

import (
	"subway/server/db"
	"subway/server/response"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

// ShowplayerDetails show all the information of the player
func ShowPlayerDetailsService(ctx *gin.Context, playerId string) {
	var playerDetails response.PlayerDetails

	var powerups []response.PowerUp

	// Extracting player details
	query1 := `SELECT p.p_id,p.p_name,p.email,p.high_score,p.total_distance,pc.coins
		FROM players as p
		JOIN player_coins as pc
		ON p.p_id=pc.p_id
		WHERE p.p_id=?`

	//Extracting player power ups
	query2 := `SELECT pu.power_up_name,ppu.quantity
				FROM players as p
				JOIN player_coins as pc
				ON p.p_id=pc.p_id
				JOIN player_power_ups as ppu
				ON ppu.p_id=p.p_id
				JOIN power_ups as pu
				ON pu.power_up_id=ppu.power_up_id
				WHERE p.p_id=?`

	err := db.RawQuery(query2, &powerups, playerId)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	playerDetails = db.Fun(query1, playerId)
	playerDetails.PowerUps = powerups

	response.ShowResponse("Success", utils.SUCCESS, "Player Details fetched sucessfully", playerDetails, ctx)

}
