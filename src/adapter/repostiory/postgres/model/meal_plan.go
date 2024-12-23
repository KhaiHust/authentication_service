package model

import "time"

type MealPlanModel struct {
	BaseModel
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	UserID      int64     `gorm:"column:user_id"`
	Schedule    time.Time `gorm:"column:schedule"`
	Status      string    `gorm:"column:status"`
}

func (MealPlanModel) TableName() string {
	return "meal_plans"
}
