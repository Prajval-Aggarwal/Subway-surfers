package avatar

import (
	"subway/server/db"
	"subway/server/model"
	"subway/server/request"
	"subway/server/response"
	"subway/server/utils"

	"github.com/gin-gonic/gin"
)

// ShowAvatarService shows the all the avatars available in the db
func ShowAvatarService(ctx *gin.Context) {
	var avatarList [2]response.AvatarResponse

	//Extracting locked avatars
	var avatarDetails []model.Avatar
	query := "select * from avatars where status='locked'"
	db.RawQuery(query, &avatarDetails)
	avatarList[0].Status = "locked"
	avatarList[0].Ava = avatarDetails

	//Extracting Unlocked avatars
	query = "select * from avatars where status='Unlocked'"
	avatarDetails = nil
	db.RawQuery(query, &avatarDetails)
	avatarList[1].Status = "Unlocked"
	avatarList[1].Ava = avatarDetails

	response.ShowResponse("Success", utils.SUCCESS, "Avatar Details fetched sucessfully", avatarList, ctx)

}

// UnlockAvatar unlocks the avatar for the player after the player meets the requirements to unlock the avatar
func UnlockAvtar(ctx *gin.Context, playerId string, playerPoints int64) {
	var avatars []model.Avatar

	//Selecting the points required to unlock the avatar
	query := "SELECT * FROM avatars WHERE status='locked' AND points_required < ?"
	err := db.RawQuery(query, &avatars, playerPoints)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

	//Unlocking the avatars
	for _, avatar := range avatars {
		avatar.Status = "Unlocked"
		playerAvatar := model.PlayerAvatar{
			P_Id:     playerId,
			AvatarId: avatar.AvatarId,
		}
		//use update record instead of create record
		//check it once again

		if !db.RecordExist("player_avatars", "avatar_id", avatar.AvatarId) {
			err = db.CreateRecord(playerAvatar)
			if err != nil {
				response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
				return
			}
		}
		err = db.UpdateRecord(&avatar, avatar.AvatarId, "avatar_id").Error
		if err != nil {
			response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
			return
		}

	}

}

// UpdateAvatarService updates the current avatar for th player
func UpdateAvatarService(ctx *gin.Context, playerId string, avatarRequest request.UpdateAvatarRequest) {

	if avatarRequest.AvatarId == "" {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Provide avatar id")
		return
	}
	// Change the current avatar of the player and check whether that avatar is unlocked or not
	var exists bool
	var playerDetails model.Player
	query := "SELECT EXISTS (SELECT * FROM player_avatars WHERE p_id=? AND avatar_id=?)"
	err := db.RawQuery(query, &exists, playerId, avatarRequest.AvatarId)
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	if !exists {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, "Avatar is not unlocked yet")
		return
	}

	err = db.FindById(&playerDetails, playerId, "p_id")
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}
	playerDetails.CurrAvatar = avatarRequest.AvatarId
	err = db.UpdateRecord(&playerDetails, playerId, "p_id").Error
	if err != nil {
		response.ErrorResponse(ctx, utils.BAD_REQUEST, err.Error())
		return
	}

}
