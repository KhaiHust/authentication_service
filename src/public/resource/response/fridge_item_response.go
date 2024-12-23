package response

import "github.com/KhaiHust/authen_service/core/entity"

type FridgeItemResponse struct {
	ID          int64         `json:"id"`
	Node        string        `json:"node"`
	ExpiredDate int64         `json:"expired_date"`
	Quantity    int           `json:"quantity"`
	FoodID      int64         `json:"food_id"`
	CreatedBy   int64         `json:"created_by"`
	CreatedAt   int64         `json:"created_at"`
	UpdatedAt   int64         `json:"updated_at"`
	Food        *FoodResponse `json:"food,omitempty"`
}

func FromEntityToFridgeItemResponse(entity *entity.FridgeItemEntity) *FridgeItemResponse {
	if entity == nil {
		return nil
	}
	var food *FoodResponse
	if entity.Food != nil {
		food = FromEntityToFoodResponse(entity.Food)
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
		Food:        food,
	}
}
func FromListEntityToFridgeItemResponse(entities []*entity.FridgeItemEntity) []*FridgeItemResponse {
	var responses []*FridgeItemResponse
	for _, entity := range entities {
		responses = append(responses, FromEntityToFridgeItemResponse(entity))
	}
	return responses
}
