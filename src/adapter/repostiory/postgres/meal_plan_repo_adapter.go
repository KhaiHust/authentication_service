package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/specification"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type MealPlanRepoAdapter struct {
	base
}

func (m MealPlanRepoAdapter) GetMealPlan(ctx context.Context, params *dto.MealPlanParams) ([]*entity.MealPlanEntity, error) {
	var mpModels []*model.MealPlanModel
	rawQuery, args := specification.BuildGetMealPlanSpecification(params)
	if err := m.db.WithContext(ctx).
		Raw("SELECT * FROM meal_plans "+rawQuery, args).Find(mpModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListMealPlanEntity(mpModels), nil
}

func (m MealPlanRepoAdapter) DeleteMealPlanByID(ctx context.Context, tx *gorm.DB, mealPlanID int64) error {
	if err := tx.WithContext(ctx).Where("id = ?", mealPlanID).Delete(&model.MealPlanModel{}).Error; err != nil {
		return err
	}
	return nil
}

func (m MealPlanRepoAdapter) UpdateMealPlan(ctx context.Context, tx *gorm.DB, mealPlanID int64, mpEntity *entity.MealPlanEntity) (*entity.MealPlanEntity, error) {
	mpModel := mapper.ToMealPlanModel(mpEntity)
	if err := tx.WithContext(ctx).Model(&model.MealPlanModel{}).Where("id = ?", mealPlanID).Updates(mpModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToMealPlanEntity(mpModel), nil
}

func (m MealPlanRepoAdapter) GetMealPlanByUserIDAndID(ctx context.Context, userID, mealPlanID int64) (*entity.MealPlanEntity, error) {
	var mpModel model.MealPlanModel
	if err := m.db.WithContext(ctx).Where("id = ? AND user_id = ?", mealPlanID, userID).First(&mpModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToMealPlanEntity(&mpModel), nil
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
