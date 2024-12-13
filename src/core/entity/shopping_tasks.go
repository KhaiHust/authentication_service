package entity

type ShoppingTaskEntity struct {
	BaseEntity
	ShoppingListID int64
	FoodName       string
	Quantity       string
	Status         string
}
