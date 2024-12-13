package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IShoppingListPort interface {
	CreateNewShoppingList(ctx context.Context, tx *gorm.DB, shoppingList *entity.ShoppingListEntity) (*entity.ShoppingListEntity, error)
	GetShoppingListByID(ctx context.Context, id int64) (*entity.ShoppingListEntity, error)
	UpdateShoppingListByID(ctx context.Context, tx *gorm.DB, shoppingList *entity.ShoppingListEntity) (*entity.ShoppingListEntity, error)
	DeleteShoppingListByID(ctx context.Context, tx *gorm.DB, id int64) error
}
