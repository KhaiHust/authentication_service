package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"time"
)

func ToMealPlanModel(mealPlanEntity *entity.MealPlanEntity) *model.MealPlanModel {
	if mealPlanEntity == nil {
		return nil
	}

	mealPlanModel := &model.MealPlanModel{
		BaseModel:   model.BaseModel{ID: mealPlanEntity.ID},
		Name:        mealPlanEntity.Name,
		Description: mealPlanEntity.Description,
		UserID:      mealPlanEntity.UserID,
		Schedule:    time.Unix(mealPlanEntity.Schedule, 0),
		Status:      mealPlanEntity.Status,
	}
	return mealPlanModel
}
func ToMealPlanEntity(mealPlanModel *model.MealPlanModel) *entity.MealPlanEntity {
	if mealPlanModel == nil {
		return nil
	}
	mealPlanEntity := &entity.MealPlanEntity{
		BaseEntity: entity.BaseEntity{ID: mealPlanModel.ID,
			CreatedAt: mealPlanModel.CreatedAt.Unix(),
			UpdatedAt: mealPlanModel.UpdatedAt.Unix()},
		Name:        mealPlanModel.Name,
		Description: mealPlanModel.Description,
		UserID:      mealPlanModel.UserID,
		Schedule:    mealPlanModel.Schedule.Unix(),
		Status:      mealPlanModel.Status,
	}
	return mealPlanEntity
}
func ToListMealPlanEntity(mealPlanModels []*model.MealPlanModel) []*entity.MealPlanEntity {
	mealPlanEntities := make([]*entity.MealPlanEntity, 0)
	for _, mealPlanModel := range mealPlanModels {
		mealPlanEntities = append(mealPlanEntities, ToMealPlanEntity(mealPlanModel))
	}
	return mealPlanEntities
}
