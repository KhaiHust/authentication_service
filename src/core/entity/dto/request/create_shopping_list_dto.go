package request

type CreateShoppingListDto struct {
	ID          int64
	Name        string
	Description string
	GroupID     int64
	AssignedTo  int64
	CreatedBy   int64
	DueDate     int64
}
