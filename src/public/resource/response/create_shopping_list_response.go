package response

import "github.com/KhaiHust/authen_service/core/entity"

type CreateShoppingListResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        int64  `json:"date"`
	CreatedBy   int64  `json:"created_by"`
	AssignedTo  int64  `json:"assigned_to"`
}

func ToCreateShoppingListResponse(entity *entity.ShoppingListEntity) *CreateShoppingListResponse {
	return &CreateShoppingListResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Date:        entity.DueDate,
		CreatedBy:   entity.CreatedBy,
		AssignedTo:  entity.AssignedTo,
	}
}
