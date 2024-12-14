package model

type ShoppingTaskModel struct {
	BaseModel
	ShoppingListID int64  `gorm:"column:shopping_list_id"`
	FoodName       string `gorm:"column:food_name"`
	Quantity       string `gorm:"column:quantity"`
	Status         string `gorm:"column:status"`
}

func (ShoppingTaskModel) TableName() string {
	return "shopping_tasks"
}
