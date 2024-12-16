package entity

type UserProfileEntity struct {
	BaseEntity
	UserID         int64
	Email          string
	Name           string
	DoB            int64
	AvatarImageUrl string
}
