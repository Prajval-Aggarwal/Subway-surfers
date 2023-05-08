package dailyreward

import (
	"fmt"
	"math/rand"
	"subway/server/db"
	"subway/server/model"
	"subway/server/response"
	"subway/server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Reward struct {
	PowerUpID string
}

// ShowPlayerRewardService shows th erward of the player that is generated on that day and is not collected yet.
func ShowPlayerRewardService(ctx *gin.Context, playerId string) {
	var playerReward model.DailyReward
	query := "SELECT * FROM daily_rewards WHERE p_id=? AND status='Not Collected'"
	err := db.RawQuery(query, &playerReward, playerId)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	response.ShowResponse("Success", utils.SUCCESS, "Successfully fetched reward", &playerReward, ctx)
}

// UpdateStreak set the streak of the player to user if player misses to collect the reward
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

// GenerateReward generates random reward for all the player
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

// RewardCollectedService collects the reward generated for the player
func RewardCollectedService(ctx *gin.Context, playerId string) {
	var playerReward model.DailyReward
	var playerPowerDetail model.PlayerPowerUps
	var playerDetails model.Player
	var exists bool

	//Check if the reward is genrated or not
	err := db.FindById(&playerReward, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Reward not generated")
		return
	}

	//Check if the player has already collected the reward or not
	if playerReward.Status == "Collected" {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Player has already collected the reward")
		return
	}

	//After collecting the reward updating the player streak
	err = db.FindById(&playerDetails, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Reward not generated")
		return
	}
	playerDetails.Streak = playerDetails.Streak + 1

	err = db.UpdateRecord(&playerDetails, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	// Changing the reward status from not collected to collected
	playerReward.Status = "Collected"
	db.UpdateRecord(&playerReward, playerId, "p_id")

	// Adding the reward to players account
	//checking if the player has taken that power up in past or not
	query := "SELECT EXISTS(SELECT * FROM player_power_ups WHERE p_id=? AND power_up_id=?)"
	db.RawQuery(query, &exists, playerId, playerReward.P_Id)
	if !exists {
		fmt.Println("if worked")
		playerPowerDetail.P_ID = playerId
		playerPowerDetail.PowerUp_Id = playerReward.PowerUpId
		playerPowerDetail.Quantity = playerReward.Quantity
		err := db.CreateRecord(&playerPowerDetail)
		if err != nil {
			response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
			return
		}
	} else {

		// If the player already has that power up the increasing its value according to the reward
		query = "SELECT * FROM player_power_ups WHERE p_id=? AND power_up_id=?;"
		db.RawQuery(query, &playerPowerDetail, playerId, playerReward.PowerUpId)

		playerPowerDetail.Quantity = playerPowerDetail.Quantity + playerReward.Quantity

		query = "UPDATE player_power_ups SET quantity=? WHERE p_id=? AND power_up_id=?;"

		err = db.ExecuteQuery(query, playerPowerDetail.Quantity, playerId, playerReward.PowerUpId)
		if err != nil {
			response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
			return
		}

	}
	response.ShowResponse("Success", utils.SUCCESS, "Reward collected successfully", nil, ctx)
}
