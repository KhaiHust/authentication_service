package response

import "github.com/KhaiHust/authen_service/core/entity"

type CreateGroupResponse struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
}

func ToCreateGroupResponse(groupEntity *entity.GroupEntity) *CreateGroupResponse {
	return &CreateGroupResponse{
		ID:          groupEntity.ID,
		Name:        groupEntity.Name,
		Description: groupEntity.Description,
		CreatedAt:   groupEntity.CreatedAt,
		UpdatedAt:   groupEntity.UpdatedAt,
	}
}
