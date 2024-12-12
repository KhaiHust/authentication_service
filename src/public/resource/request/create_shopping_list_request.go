package request

type CreateShoppingListRequest struct {
	Name        string `json:"name" validate:"required,gte=3" message:"validateName" errorCode:"400057"`
	Description string `json:"description"`
	GroupID     int64  `json:"group_id"`
	AssignedTo  int64  `json:"assigned_to"`
	DueDate     int64  `json:"due_date" validate:"required,due_date_time" message:"validateDueDate" errorCode:"400058"`
}
