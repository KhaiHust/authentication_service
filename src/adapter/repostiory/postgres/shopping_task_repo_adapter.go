package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type ShoppingTaskRepoAdapter struct {
	base
}

func (s *ShoppingTaskRepoAdapter) CreateNewShoppingTasks(ctx context.Context, tx *gorm.DB, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error) {
	shoppingTaskModels := mapper.ToListShoppingTaskModel(shoppingTasks)
	if err := tx.WithContext(ctx).Create(&shoppingTaskModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListShoppingTaskEntity(shoppingTaskModels), nil
}
func NewShoppingTaskRepoAdapter(db *gorm.DB) port.IShoppingTaskPort {
	return &ShoppingTaskRepoAdapter{base{db}}
}
