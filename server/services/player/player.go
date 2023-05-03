package player

import (
	"subway/server/db"
	"subway/server/response"

	"github.com/gin-gonic/gin"
)

func ShowPlayerDetailsService(ctx *gin.Context, playerId string) {
	ctx.Header("Content-Type", "application/json")
	query1 := `SELECT p.p_id,p.p_name,p.email,p.high_score,p.total_distance,pc.coins
		FROM players as p
		JOIN player_coins as pc
		ON p.p_id=pc.p_id
		WHERE p.p_id=?`

	query2 := `SELECT pu.power_up_name,ppu.quantity
				FROM players as p
				JOIN player_coins as pc
				ON p.p_id=pc.p_id
				JOIN player_power_ups as ppu
				ON ppu.p_id=p.p_id
				JOIN power_ups as pu
				ON pu.power_up_id=ppu.power_up_id
				WHERE p.p_id=?`

	//mp := make(map[response.PlayerDetails][]response.PowerUp)
	var playerDetails response.PlayerDetails

	var powerups []response.PowerUp
	err := db.RawQuery(query2, &powerups, playerId)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	playerDetails = db.Fun(query1, playerId)
	playerDetails.PowerUps = powerups

	response.ShowResponse("Success", 200, "Player Details fetched sucessfully", playerDetails, ctx)

}
