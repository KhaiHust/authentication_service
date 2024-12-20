package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type FridgeItemRepositoryAdapter struct {
	base
}

func (f FridgeItemRepositoryAdapter) UpdateItem(ctx context.Context, tx *gorm.DB, fridgeItem *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error) {
	itemModel := mapper.ToFridgeItemModel(fridgeItem)
	if err := tx.WithContext(ctx).Model(&model.FridgeItemModel{}).Where("id = ?", fridgeItem.ID).Updates(itemModel).Error; err != nil {
		return nil, err
	}

	return mapper.ToFridgeItemEntity(itemModel), nil
}

func (f FridgeItemRepositoryAdapter) GetItemByIDAndCreatedBy(ctx context.Context, id, createdBy int64) (*entity.FridgeItemEntity, error) {
	itemModel := &model.FridgeItemModel{}
	if err := f.db.WithContext(ctx).Model(&model.FridgeItemModel{}).Where("id = ? AND created_by = ?", id, createdBy).First(itemModel).Error; err != nil {
		return nil, err
	}

	return mapper.ToFridgeItemEntity(itemModel), nil
}

func (f FridgeItemRepositoryAdapter) SaveFridgeItem(ctx context.Context, tx *gorm.DB, fridgeItem *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error) {
	itemModel := mapper.ToFridgeItemModel(fridgeItem)
	if err := tx.WithContext(ctx).Model(itemModel).Create(itemModel).Error; err != nil {
		return nil, err
	}

	return mapper.ToFridgeItemEntity(itemModel), nil
}

func NewFridgeItemRepositoryAdapter(db *gorm.DB) port.IFridgeItemPort {
	return &FridgeItemRepositoryAdapter{
		base: base{db: db},
	}
}
