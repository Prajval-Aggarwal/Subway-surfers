package dailyreward

import (
	"fmt"
	"math/rand"
	"subway/server/db"
	"subway/server/model"
	"subway/server/response"
	"time"

	"github.com/gin-gonic/gin"
)

type Reward struct {
	PowerUpID string
}

func ShowPlayerRewardService(ctx *gin.Context, playerId string) {
	var playerReward model.DailyReward
	query := "SELECT * FROM daily_rewards WHERE p_id=? AND status='Not Collected'"
	err := db.RawQuery(query, &playerReward, playerId)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	response.ShowResponse("Success", 200, "Successfully fetched reward", &playerReward, ctx)
}

func UpdateStreak() error {
	var playersReward []model.DailyReward
	query := "SELECT * FROM daily_rewards WHERE status='Not Collected'"
	db.RawQuery(query, &playersReward)

	for _, p := range playersReward {
		var playerDetails model.Player
		err := db.FindById(&playerDetails, p.P_Id, "p_id")
		if err != nil {
			return err
		}
		playerDetails.Streak = 0
		err = db.UpdateRecord(&playerDetails, p.P_Id, "p_id").Error
		if err != nil {
			return err
		}
	}
	return nil

}

func GenerateReward() error {
	query := "TRUNCATE TABLE daily_rewards"
	db.ExecuteQuery(query)
	var players []model.Player
	query = "SELECT p_id FROM players"
	db.RawQuery(query, &players)

	for _, p := range players {

		var dailyReward model.DailyReward
		dailyReward.P_Id = p.P_ID
		now := time.Now().Truncate(time.Hour)
		dailyReward.CreatedAt = now
		dailyReward.Status = "Not Collected"
		dailyReward.Quantity = int64(rand.Intn(10-1) + 1)
		reward := SelectRand()
		dailyReward.PowerUpId = reward.PowerUpID
		err := db.CreateRecord(&dailyReward)
		if err != nil {
			return err
		}
	}
	return nil

}

func SelectRand() *Reward {

	rewrd := &Reward{}
	query := "SELECT power_up_id FROM power_ups ORDER BY RANDOM()  LIMIT 1;"
	err := db.RawQuery(query, rewrd)
	if err != nil {
		fmt.Println("error is:", err)
		return nil
	}

	return rewrd
}

func RewardCollectedService(ctx *gin.Context, playerId string) {
	var playerReward model.DailyReward
	var playerPowerDetail model.PlayerPowerUps
	var playerDetails model.Player
	var exists bool
	err := db.FindById(&playerReward, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 400, "Reward not generated")
		return
	}
	if playerReward.Status == "Collected" {
		response.ErrorResponse(ctx, 400, "Player has already collected the reward")
		return
	}
	//update streak
	err = db.FindById(&playerDetails, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 400, "Reward not generated")
		return
	}
	playerDetails.Streak = playerDetails.Streak + 1

	err = db.UpdateRecord(&playerDetails, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

	//change reward status from not collected to collected
	playerReward.Status = "Collected"
	db.UpdateRecord(&playerReward, playerId, "p_id")

	query := "SELECT EXISTS(SELECT * FROM player_power_ups WHERE p_id=? AND power_up_id=?)"
	db.RawQuery(query, &exists, playerId, playerReward.P_Id)
	if !exists {
		fmt.Println("if worked")
		playerPowerDetail.P_ID = playerId
		playerPowerDetail.PowerUp_Id = playerReward.PowerUpId
		playerPowerDetail.Quantity = playerReward.Quantity
		db.CreateRecord(&playerPowerDetail)
	} else {

		query = "SELECT * FROM player_power_ups WHERE p_id=? AND power_up_id=?;"
		db.RawQuery(query, &playerPowerDetail, playerId, playerReward.PowerUpId)

		fmt.Println("else worked")
		fmt.Println("before player power up quantity is", playerPowerDetail.Quantity)
		playerPowerDetail.Quantity = playerPowerDetail.Quantity + playerReward.Quantity
		fmt.Println("player power up quantity is", playerPowerDetail.Quantity)

		query = "UPDATE player_power_ups SET quantity=? WHERE p_id=? AND power_up_id=?;"

		err = db.ExecuteQuery(query, playerPowerDetail.Quantity, playerId, playerReward.PowerUpId)
		if err != nil {
			response.ErrorResponse(ctx, 400, err.Error())
			return
		}

	}
}
