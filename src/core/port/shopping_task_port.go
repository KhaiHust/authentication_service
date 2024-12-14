package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IShoppingTaskPort interface {
	CreateNewShoppingTasks(ctx context.Context, tx *gorm.DB, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error)
	GetShoppingTasksByShoppingListID(ctx context.Context, shoppingListID int64) ([]*entity.ShoppingTaskEntity, error)
}
