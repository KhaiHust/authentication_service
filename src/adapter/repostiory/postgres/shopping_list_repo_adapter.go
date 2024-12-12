package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type ShoppingListRepoAdapter struct {
	base
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
