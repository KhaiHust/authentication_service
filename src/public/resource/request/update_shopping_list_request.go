package request

import "github.com/KhaiHust/authen_service/core/entity/dto/request"

type UpdateShoppingListRequest struct {
	Name        string `json:"new_name"`
	Description string `json:"new_description"`
	AssignedTo  int64  `json:"new_assigned_to"`
	DueDate     int64  `json:"new_due_date"`
}

func ToUpdateShoppingListDto(req *UpdateShoppingListRequest) *request.CreateShoppingListDto {
	if req == nil {
		return nil
	}
	return &request.CreateShoppingListDto{
		Name:        req.Name,
		Description: req.Description,
		AssignedTo:  req.AssignedTo,
		DueDate:     req.DueDate,
	}
}
