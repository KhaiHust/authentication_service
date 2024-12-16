package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type ShoppingListGroupRepoAdapter struct {
	base
}

func (s ShoppingListGroupRepoAdapter) GetShoppingListGroupByShoppingListID(ctx context.Context, shoppingListID int64) (*entity.ShoppingListGroupEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (s ShoppingListGroupRepoAdapter) CreateNewShoppingListGroup(ctx context.Context, tx *gorm.DB, shoppingListGroup *entity.ShoppingListGroupEntity) (*entity.ShoppingListGroupEntity, error) {
	sLGModel := mapper.ToShoppingListGroupModel(shoppingListGroup)
	if err := tx.WithContext(ctx).Model(&model.ShoppingListGroupModel{}).Create(sLGModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToShoppingListGroupEntity(sLGModel), nil
}

func NewShoppingListGroupRepoAdapter(db *gorm.DB) port.IShoppingListGroupPort {
	return &ShoppingListGroupRepoAdapter{
		base: base{db: db},
	}
}
