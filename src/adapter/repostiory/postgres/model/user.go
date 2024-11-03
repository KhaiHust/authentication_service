package model

type UserModel struct {
	BaseModel
	Name       string  `gorm:"column:name" sql:"not null"`
	Email      string  `gorm:"column:email" sql:"not null, unique"`
	Password   string  `gorm:"column:password_hash" sql:"not null"`
	DeviceID   *string `gorm:"column:device_id"`
	Language   *string `gorm:"column:language"`
	Timezone   *string `gorm:"column:timezone"`
	IsActive   bool    `gorm:"column:is_active" sql:"default:true"`
	IsVerified bool    `gorm:"column:is_verified" sql:"default:false"`
}

func (*UserModel) TableName() string {
	return "users"
}
