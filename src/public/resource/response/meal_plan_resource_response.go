package response

import "github.com/KhaiHust/authen_service/core/entity"

type MealPlanResourceResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Schedule    int64           `json:"schedule"`
	Status      string          `json:"status"`
	FoodIDs     []int64         `json:"food_ids"`
	Foods       []*FoodResponse `json:"foods,omitempty"`
}

func FromEntityToMealPlanResourceResponse(entity *entity.MealPlanEntity) *MealPlanResourceResponse {
	if entity == nil {
		return nil
	}
	return &MealPlanResourceResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Schedule:    entity.Schedule,
		Status:      entity.Status,
		FoodIDs:     entity.FoodIDs,
	}
}
func FromEntitiesToMealPlanResourceResponses(entities []*entity.MealPlanEntity) []*MealPlanResourceResponse {
	if entities == nil {
		return nil
	}
	responses := make([]*MealPlanResourceResponse, 0)
	for _, entity := range entities {
		responses = append(responses, FromEntityToMealPlanResourceResponse(entity))
	}
	return responses
}
