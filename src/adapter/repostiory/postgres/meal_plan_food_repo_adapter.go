package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type MealPlanFoodRepoAdapter struct {
	base
}

func (m MealPlanFoodRepoAdapter) DeleteListMealPlanFood(ctx context.Context, tx *gorm.DB, mealPlanID int64) error {
	if err := tx.WithContext(ctx).Where("meal_plan_id = ?", mealPlanIDs).Delete(&model.MealPlanFoodModel{}).Error; err != nil {
		return err
	}
	return nil
}

func (m MealPlanFoodRepoAdapter) SaveListMealPlanFood(ctx context.Context, tx *gorm.DB, mpFEntities []*entity.MealPlanFoodEntity) ([]*entity.MealPlanFoodEntity, error) {
	mpFModels := mapper.ToListMealPlanFoodModel(mpFEntities)
	if err := tx.WithContext(ctx).Model(&model.MealPlanFoodModel{}).Create(mpFModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListMealPlanFoodEntity(mpFModels), nil
}

func (m MealPlanFoodRepoAdapter) SaveMealPlanFood(ctx context.Context, tx *gorm.DB, mpFEntity *entity.MealPlanFoodEntity) (*entity.MealPlanFoodEntity, error) {
	mpFModel := mapper.ToMealPlanFoodModel(mpFEntity)
	if err := tx.WithContext(ctx).Model(&model.MealPlanFoodModel{}).Create(mpFModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToMealPlanFoodEntity(mpFModel), nil
}

func NewMealPlanFoodRepoAdapter(db *gorm.DB) port.IMealPlanFoodPort {
	return &MealPlanFoodRepoAdapter{
		base: base{db: db},
	}
}
