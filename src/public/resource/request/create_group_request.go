package request

import "github.com/KhaiHust/authen_service/core/entity/dto/request"

type CreateGroupRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func ToCreateGroupRequestDto(req *CreateGroupRequest) *request.CreateGroupDTO {
	return &request.CreateGroupDTO{
		Name:        req.Name,
		Description: req.Description,
	}
}
