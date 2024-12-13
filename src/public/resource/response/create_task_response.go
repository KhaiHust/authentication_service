package response

import "github.com/KhaiHust/authen_service/core/entity"

type CreateTaskResponse struct {
	ShoppingListID int64   `json:"shopping_list_id"`
	Tasks          []Tasks `json:"tasks"`
}
type Tasks struct {
	ID        int64  `json:"id"`
	FoodName  string `json:"food_name"`
	Quantity  string `json:"quantity"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func ToCreateTaskResponse(shoppingListID int64, tasks []*entity.ShoppingTaskEntity) *CreateTaskResponse {
	var response CreateTaskResponse
	response.ShoppingListID = shoppingListID
	for _, task := range tasks {
		response.Tasks = append(response.Tasks, Tasks{
			ID:        task.ID,
			FoodName:  task.FoodName,
			Quantity:  task.Quantity,
			Status:    task.Status,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		})
	}
	return &response
}
