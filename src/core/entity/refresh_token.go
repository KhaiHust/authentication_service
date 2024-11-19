package entity

type RefreshTokenEntity struct {
	BaseEntity
	UserId       int64
	RefreshToken string
	ExpiredAt    int64
	IpAddress    string
	UserAgent    string
}
