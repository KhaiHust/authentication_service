package response

import "github.com/KhaiHust/authen_service/core/entity"

type RegisterUserResponse struct {
	Email      string `json:"email,omitempty"`
	Name       string `json:"name,omitempty"`
	IsVerified bool   `json:"is_verified"`
	IsActive   bool   `json:"is_active"`
}

func FromEntityToRegisterUserResponse(entity *entity.UserEntity) *RegisterUserResponse {
	if entity == nil {
		return nil
	}
	return &RegisterUserResponse{
		Email:      entity.Email,
		Name:       entity.Name,
		IsVerified: entity.IsVerified,
		IsActive:   entity.IsActive,
	}
}
