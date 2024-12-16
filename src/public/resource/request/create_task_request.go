package request

import "github.com/KhaiHust/authen_service/core/entity"

type CreateShoppingTaskRequest struct {
	Tasks []Task `json:"tasks" validate:"required" message:"tasks is required"`
}
type Task struct {
	FoodName string `json:"food_name" validate:"required" message:"food_name is required"`
	Quantity string `json:"quantity" validate:"required" message:"quantity is required"`
}

func ToCreateTasksEntity(request CreateShoppingTaskRequest) []*entity.ShoppingTaskEntity {
	var entities []*entity.ShoppingTaskEntity
	for _, task := range request.Tasks {
		entities = append(entities, &entity.ShoppingTaskEntity{
			FoodName: task.FoodName,
			Quantity: task.Quantity,
		})
	}
	return entities
}
