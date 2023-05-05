package avatar

import (
	"fmt"
	"subway/server/db"
	"subway/server/model"
	"subway/server/request"
	"subway/server/response"

	"github.com/gin-gonic/gin"
)

func ShowAvatarService(ctx *gin.Context) {
	var avatarList [2]response.AvatarResponse
	//show locked and unlocked avatar list
	var avatarDetails []model.Avatar
	query := "select * from avatars where status='locked'"
	db.RawQuery(query, &avatarDetails)
	avatarList[0].Status = "locked"
	avatarList[0].Ava = avatarDetails
	fmt.Println("pehla: ", avatarList[0])

	query = "select * from avatars where status='Unlocked'"
	avatarDetails = nil
	db.RawQuery(query, &avatarDetails)
	avatarList[1].Status = "Unlocked"
	avatarList[1].Ava = avatarDetails
	fmt.Println("dusra ka phela", avatarList[0])
	fmt.Println("dusra: ", avatarList[1])

	response.ShowResponse("Success", 200, "Avatar Details fetched sucessfully", avatarList, ctx)

}

func UnlockAvtar(ctx *gin.Context, playerId string, playerPoints int64) {
	var avatars []model.Avatar
	fmt.Println("playerasda", playerPoints)
	query := "SELECT * FROM avatars WHERE points_required < ?"
	err := db.RawQuery(query, &avatars, playerPoints)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	for _, avatar := range avatars {
		avatar.Status = "Unlocked"
		var playerAvatar model.PlayerAvatar
		playerAvatar.P_Id = playerId
		playerAvatar.AvatarId = avatar.AvatarId

		//use update record instead of create record

		if !db.RecordExist("player_avatars", "avatar_id", avatar.AvatarId) {
			err = db.CreateRecord(playerAvatar)
			if err != nil {
				response.ErrorResponse(ctx, 400, err.Error())
				return
			}
		}
		err = db.UpdateRecord(&avatar, avatar.AvatarId, "avatar_id").Error
		if err != nil {
			response.ErrorResponse(ctx, 400, err.Error())
			return
		}

	}

}

func UpdateAvatarService(ctx *gin.Context, playerId string, avatarRequest request.UpdateAvatarRequest) {
	//chnage the current avatar of the user and chack whtether taht avar tar is present in playerAvatar list or not
	var exists bool
	query := "SELECT EXISTS (SELECT * FROM player_avatars WHERE p_id=? AND avatar_id=?)"
	err := db.RawQuery(query, &exists, playerId, avatarRequest.AvatarId)
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	if !exists {
		response.ErrorResponse(ctx, 400, "Avatar is not unlocked yet")
		return
	}
	var playerDetails model.Player
	err = db.FindById(&playerDetails, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}
	playerDetails.CurrAvatar = avatarRequest.AvatarId
	err = db.UpdateRecord(&playerDetails, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, 400, err.Error())
		return
	}

}
