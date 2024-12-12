package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IShoppingListGroupPort interface {
	CreateNewShoppingListGroup(ctx context.Context, tx *gorm.DB, shoppingListGroup *entity.ShoppingListGroupEntity) (*entity.ShoppingListGroupEntity, error)
}
