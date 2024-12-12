package model

type ShoppingListModel struct {
	BaseModel
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	CreatedBy   int64  `gorm:"column:created_by"`
	AssignedTo  int64  `gorm:"column:assigned_to"`
}

func (m *ShoppingListModel) TableName() string {
	return "shopping_lists"
}
