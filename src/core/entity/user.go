package entity

type UserEntity struct {
	BaseEntity
	Name       string
	Email      string
	Password   string
	DeviceID   *string
	Language   *string
	Timezone   *string
	IsActive   bool
	IsVerified bool
}
