package model

type ShoppingListGroup struct {
	BaseModel
	ShoppingListID int64 `gorm:"column:shopping_list_id"`
	GroupID        int64 `gorm:"column:group_id"`
}

func (m *ShoppingListGroup) TableName() string {
	return "shopping_list_groups"
}
