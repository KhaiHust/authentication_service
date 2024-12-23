package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type MealPlanRepoAdapter struct {
	base
}

func (m MealPlanRepoAdapter) SaveNewMealPlan(ctx context.Context, tx *gorm.DB, mpEntity *entity.MealPlanEntity) (*entity.MealPlanEntity, error) {
	mpModel := mapper.ToMealPlanModel(mpEntity)
	if err := tx.WithContext(ctx).Model(&model.MealPlanModel{}).Create(mpModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToMealPlanEntity(mpModel), nil
}

func NewMealPlanRepoAdapter(db *gorm.DB) port.IMealPlanPort {
	return &MealPlanRepoAdapter{
		base: base{db: db},
	}
}
