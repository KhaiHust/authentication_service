package request

import "github.com/KhaiHust/authen_service/core/entity"

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"email" message:"Email is required" errorCode:"400026"`
	Password string `json:"password" validate:"gt=6, lt=20" message:"Password is required and must be between 7 and 19 characters" errorCode:"400027"`
	Name     string `json:"name" validate:"gt=3, lt=30" message:"Name is required and must be between 4 and 29 characters" errorCode:"400028"`
	Language string `json:"language" default:"vn"`
	TimeZone string `json:"timeZone" default:"UTC"`
	DeviceID string `json:"deviceId" validate:"required"`
}

func (r *RegisterUserRequest) ToEntity() *entity.UserEntity {
	return &entity.UserEntity{
		Email:    r.Email,
		Password: r.Password,
		Name:     r.Name,
		Language: &r.Language,
		Timezone: &r.TimeZone,
		DeviceID: &r.DeviceID,
	}
}
