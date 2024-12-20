package entity

type FridgeItemEntity struct {
	BaseEntity
	ExpiredDate int64
	Quantity    int
	FoodID      int64
	CreatedBy   int64
	Note        string
}
