package model

import "time"

type ShoppingListModel struct {
	BaseModel
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	DueDate     time.Time `gorm:"column:due_date"`
	CreatedBy   int64     `gorm:"column:created_by"`
	AssignedTo  int64     `gorm:"column:assigned_to"`
}

func (m *ShoppingListModel) TableName() string {
	return "shopping_lists"
}
