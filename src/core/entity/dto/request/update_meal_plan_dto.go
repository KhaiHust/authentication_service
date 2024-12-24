package request

type UpdateMealPlanDTO struct {
	Name        *string
	Description *string
	Schedule    *int64
	Status      *string
	FoodIDs     []int64
}
