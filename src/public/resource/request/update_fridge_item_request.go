package request

import request2 "github.com/KhaiHust/authen_service/core/entity/dto/request"

type UpdateFridgeItemRequest struct {
	Quantity    *int    `json:"quantity"`
	ExpiredDate *int64  `json:"expired_date"`
	Node        *string `json:"node"`
	FoodID      *int64  `json:"food_id"`
}

func ToUpdateFridgeItemDto(request *UpdateFridgeItemRequest) *request2.UpdateFridgeItemDto {
	if request == nil {
		return nil
	}
	return &request2.UpdateFridgeItemDto{
		ExpiredDate: request.ExpiredDate,
		Quantity:    request.Quantity,
		Node:        request.Node,
		FoodID:      request.FoodID,
	}
}
