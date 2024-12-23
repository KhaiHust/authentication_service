package request

type UpdateFridgeItemDto struct {
	Quantity    *int
	ExpiredDate *int64
	Node        *string
	FoodID      *int64
}
