package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func ToMealPlanFoodModel(meal *entity.MealPlanFoodEntity) *model.MealPlanFoodModel {
	if meal == nil {
		return nil
	}
	return &model.MealPlanFoodModel{
		BaseModel:  model.BaseModel{ID: meal.ID},
		MealPlanID: meal.MealPlanID,
		FoodID:     meal.FoodID,
	}
}
func ToMealPlanFoodEntity(meal *model.MealPlanFoodModel) *entity.MealPlanFoodEntity {
	if meal == nil {
		return nil
	}
	return &entity.MealPlanFoodEntity{
		BaseEntity: entity.BaseEntity{ID: meal.ID,
			CreatedAt: meal.CreatedAt.Unix(),
			UpdatedAt: meal.UpdatedAt.Unix()},
		MealPlanID: meal.MealPlanID,
		FoodID:     meal.FoodID,
	}
}
func ToListMealPlanFoodEntity(meals []*model.MealPlanFoodModel) []*entity.MealPlanFoodEntity {
	if meals == nil {
		return nil
	}
	var result []*entity.MealPlanFoodEntity
	for _, meal := range meals {
		result = append(result, ToMealPlanFoodEntity(meal))
	}
	return result
}
func ToListMealPlanFoodModel(meals []*entity.MealPlanFoodEntity) []*model.MealPlanFoodModel {
	if meals == nil {
		return nil
	}
	var result []*model.MealPlanFoodModel
	for _, meal := range meals {
		result = append(result, ToMealPlanFoodModel(meal))
	}
	return result
}
