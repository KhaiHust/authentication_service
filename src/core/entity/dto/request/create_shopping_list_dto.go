package request

type CreateShoppingListDto struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	GroupID     int64  `json:"group_id,omitempty"`
	AssignedTo  int64  `json:"assigned_to"`
	CreatedBy   int64  `json:"created_by"`
	DueDate     int64  `json:"due_date"`
}
