package model

import "time"

type UserProfileModel struct {
	BaseModel
	UserID         int64     `gorm:"column:user_id"`
	Email          string    `gorm:"column:email"`
	Name           string    `gorm:"column:name"`
	DoB            time.Time `gorm:"column:dob"`
	AvatarImageUrl string    `gorm:"column:avatar_image_url"`
}

func (UserProfileModel) TableName() string {
	return "user_profiles"
}
