package entity

type ShoppingListEntity struct {
	BaseEntity
	Name        string
	Description string
	CreatedBy   int64
	AssignedTo  int64
}
