package model

import "time"

type RefreshTokenModel struct {
	BaseModel
	UserId       int64     `gorm:"column:user_id;not null"`
	RefreshToken string    `gorm:"column:refresh_token;not null"`
	ExpiredAt    time.Time `gorm:"column:expired_at;not null"`
	IpAddress    string    `gorm:"column:ip_address"`
	UserAgent    string    `gorm:"column:user_agent"`
}

func (RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}
