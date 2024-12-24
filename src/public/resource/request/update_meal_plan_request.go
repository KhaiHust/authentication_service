package request

import "github.com/KhaiHust/authen_service/core/entity/dto/request"

type UpdateMealPlanRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Schedule    *int64  `json:"schedule"`
	Status      *string `json:"status"`
	FoodIDs     []int64 `json:"food_ids"`
}

func FromReqToUpdateMealPlanDto(req *UpdateMealPlanRequest) *request.UpdateMealPlanDTO {
	if req == nil {
		return nil
	}
	return &request.UpdateMealPlanDTO{
		Name:        req.Name,
		Description: req.Description,
		Schedule:    req.Schedule,
		Status:      req.Status,
		FoodIDs:     req.FoodIDs,
	}
}
