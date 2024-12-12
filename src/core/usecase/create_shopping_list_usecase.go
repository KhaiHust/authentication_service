package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
)

type ICreateShoppingListUseCase interface {
	CreateNewShoppingList(ctx context.Context, shoppingList *entity.ShoppingListEntity) (*entity.ShoppingListEntity, error)
}
