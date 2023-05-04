package model

type Avatar struct {
	AvatarId       string `json:"avatarId" gorm:"default:uuid_generate_v4();primaryKey"`
	AvatarName     string
	PointsRequired int64
	Status         string `json:"status" gorm:"default:locked"`
}
