package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IShoppingListPort interface {
	CreateNewShoppingList(ctx context.Context, tx *gorm.DB, shoppingList *entity.ShoppingListEntity) (*entity.ShoppingListEntity, error)
}
