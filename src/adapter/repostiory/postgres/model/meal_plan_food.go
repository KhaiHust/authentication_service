package model

type MealPlanFoodModel struct {
	BaseModel
	MealPlanID int64 `gorm:"column:meal_plan_id"`
	FoodID     int64 `gorm:"column:food_id"`
}

func (MealPlanFoodModel) TableName() string {
	return "meal_plan_foods"
}
