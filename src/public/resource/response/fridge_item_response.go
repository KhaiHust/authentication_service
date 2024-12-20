package response

import "github.com/KhaiHust/authen_service/core/entity"

type FridgeItemResponse struct {
	ID          int64  `json:"id"`
	Node        string `json:"node"`
	ExpiredDate int64  `json:"expired_date"`
	Quantity    int    `json:"quantity"`
	FoodID      int64  `json:"food_id"`
	CreatedBy   int64  `json:"created_by"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func FromEntityToFridgeItemResponse(entity *entity.FridgeItemEntity) *FridgeItemResponse {
	if entity == nil {
		return nil
	}
	return &FridgeItemResponse{
		ID:          entity.ID,
		ExpiredDate: entity.ExpiredDate,
		Quantity:    entity.Quantity,
		FoodID:      entity.FoodID,
		CreatedBy:   entity.CreatedBy,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		Node:        entity.Note,
	}
}
