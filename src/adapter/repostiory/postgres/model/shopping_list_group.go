package model

type ShoppingListGroupModel struct {
	BaseModel
	ShoppingListID int64 `gorm:"column:shopping_list_id"`
	GroupID        int64 `gorm:"column:group_id"`
}

func (m *ShoppingListGroupModel) TableName() string {
	return "shopping_list_groups"
}
