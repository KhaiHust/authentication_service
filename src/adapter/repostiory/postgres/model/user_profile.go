package model

import "time"

type UserProfileModel struct {
	BaseModel
	UserID         int64     `json:"user_id"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	DoB            time.Time `json:"dob"`
	AvatarImageUrl string    `json:"avatar_image_url"`
}

func (UserProfileModel) TableName() string {
	return "user_profiles"
}
