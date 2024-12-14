package request

type UpdateTaskDto struct {
	FoodName *string `json:"food_name"`
	Quantity *string `json:"quantity"`
	Status   *string `json:"status"`
}
