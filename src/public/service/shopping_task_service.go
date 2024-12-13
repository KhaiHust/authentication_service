package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IShoppingTaskService interface {
	CreateNewShoppingTask(ctx context.Context, userID, shoppingListId int64, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error)
}
type ShoppingTaskService struct {
	createShoppingTaskUsecase usecase.ICreateShoppingTaskUsecase
}

func (s ShoppingTaskService) CreateNewShoppingTask(ctx context.Context, userID, shoppingListId int64, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error) {
	return s.createShoppingTaskUsecase.CreateNewShoppingTask(ctx, userID, shoppingListId, shoppingTasks)
}

func NewShoppingTaskService(createShoppingTaskUsecase usecase.ICreateShoppingTaskUsecase) IShoppingTaskService {
	return &ShoppingTaskService{createShoppingTaskUsecase}
}
