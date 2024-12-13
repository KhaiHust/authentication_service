package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type ShoppingListRepoAdapter struct {
	base
}

func (s ShoppingListRepoAdapter) DeleteShoppingListByID(ctx context.Context, tx *gorm.DB, id int64) error {
	if err := tx.WithContext(ctx).Model(&model.ShoppingListModel{}).Where("id = ?", id).Delete(&model.ShoppingListModel{}).Error; err != nil {
		return err
	}
	return nil
}

func (s ShoppingListRepoAdapter) UpdateShoppingListByID(ctx context.Context, tx *gorm.DB, shoppingList *entity.ShoppingListEntity) (*entity.ShoppingListEntity, error) {
	sLModel := mapper.ToShoppingListModel(shoppingList)
	if err := tx.WithContext(ctx).Model(&model.ShoppingListModel{}).Where("id = ?", shoppingList.ID).Updates(sLModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToShoppingListEntity(sLModel), nil
}

func (s ShoppingListRepoAdapter) GetShoppingListByID(ctx context.Context, id int64) (*entity.ShoppingListEntity, error) {
	var shoppingList model.ShoppingListModel
	if err := s.db.WithContext(ctx).Model(&model.ShoppingListModel{}).Where("id = ?", id).First(&shoppingList).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf(constant.ErrShoppingListNotFound)
		}
		return nil, err
	}
	return mapper.ToShoppingListEntity(&shoppingList), nil
}

func (s ShoppingListRepoAdapter) CreateNewShoppingList(ctx context.Context, tx *gorm.DB, shoppingList *entity.ShoppingListEntity) (*entity.ShoppingListEntity, error) {
	sLModel := mapper.ToShoppingListModel(shoppingList)
	if err := tx.WithContext(ctx).Model(&model.ShoppingListModel{}).Create(sLModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToShoppingListEntity(sLModel), nil
}

func NewShoppingListRepoAdapter(db *gorm.DB) port.IShoppingListPort {
	return &ShoppingListRepoAdapter{
		base: base{db: db},
	}
}
