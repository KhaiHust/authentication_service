package request

import "github.com/KhaiHust/authen_service/core/entity"

type CreateNewMealPlanRequest struct {
	Name        string  `json:"name" validate:"required" message:"Name is required"`
	Description string  `json:"description"`
	Schedule    int64   `json:"schedule" validate:"required" message:"Schedule is required"`
	Status      string  `json:"status" validate:"required" message:"Status is required"`
	FoodIDs     []int64 `json:"food_ids"`
}

func FromCreateRequest(req *CreateNewMealPlanRequest) *entity.MealPlanEntity {
	return &entity.MealPlanEntity{
		Name:        req.Name,
		Description: req.Description,
		Schedule:    req.Schedule,
		Status:      req.Status,
		FoodIDs:     req.FoodIDs,
	}
}
