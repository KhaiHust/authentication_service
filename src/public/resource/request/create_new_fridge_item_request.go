package request

import "github.com/KhaiHust/authen_service/core/entity"

type CreateNewFridgeItemRequest struct {
	FoodID      int64  `json:"food_id" validate:"required" message:"food_id is required"`
	Quantity    int    `json:"quantity" validate:"required" message:"quantity is required"`
	ExpiredDate int64  `json:"expired_date" validate:"required" message:"expired_date is required"`
	Node        string `json:"node"`
}

func FromRequestToFridgeItemEntity(request *CreateNewFridgeItemRequest) *entity.FridgeItemEntity {
	if request == nil {
		return nil
	}
	return &entity.FridgeItemEntity{
		ExpiredDate: request.ExpiredDate,
		Quantity:    request.Quantity,
		FoodID:      request.FoodID,
		Note:        request.Node,
	}
}
