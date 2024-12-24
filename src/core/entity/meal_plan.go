package entity

type MealPlanEntity struct {
	BaseEntity
	Name        string
	Description string
	UserID      int64
	Schedule    int64
	Status      string
	FoodIDs     []int64
	Foods       []*FoodEntity
}
