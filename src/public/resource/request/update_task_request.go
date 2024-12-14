package request

import "github.com/KhaiHust/authen_service/core/entity/dto/request"

type UpdateTaskRequest struct {
	FoodName *string `json:"food_name"`
	Quantity *string `json:"quantity"`
	Status   *string `json:"status"`
}

func ToUpdateTaskDto(req *UpdateTaskRequest) *request.UpdateTaskDto {
	return &request.UpdateTaskDto{
		FoodName: req.FoodName,
		Quantity: req.Quantity,
		Status:   req.Status,
	}
}
