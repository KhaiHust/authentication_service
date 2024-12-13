package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IShoppingTaskPort interface {
	CreateNewShoppingTasks(ctx context.Context, tx *gorm.DB, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error)
}
