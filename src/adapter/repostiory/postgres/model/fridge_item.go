package model

import "time"

type FridgeItemModel struct {
	BaseModel
	Note        string    `gorm:"column:note"`
	ExpiredDate time.Time `gorm:"column:expired_date;not null"`
	Quantity    int       `gorm:"column:quantity;not null"`
	FoodID      int64     `gorm:"column:food_id;not null"`
	CreatedBy   int64     `gorm:"column:created_by;not null"`
}

func (FridgeItemModel) TableName() string {
	return "fridge_items"
}
