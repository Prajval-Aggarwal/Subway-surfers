package response

import "subway/server/model"

type AvatarResponse struct {
	Status string         `json:"status"`
	Ava    []model.Avatar `json:"avatars"`
}
