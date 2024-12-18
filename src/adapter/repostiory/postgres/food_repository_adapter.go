package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type FoodRepositoryAdapter struct {
	base
}

func (f FoodRepositoryAdapter) SaveFood(ctx context.Context, tx *gorm.DB, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error) {
	foodModel := mapper.ToFoodModel(foodEntity)
	if err := tx.WithContext(ctx).Model(&model.FoodModel{}).Create(&foodModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToFoodEntity(foodModel), nil
}

func NewFoodRepositoryAdapter(db *gorm.DB) port.IFoodPort {
	return &FoodRepositoryAdapter{
		base: base{db: db},
	}
}
