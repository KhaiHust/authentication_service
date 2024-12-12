package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IShoppingListService interface {
	CreateNewShoppingList(ctx context.Context, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error)
}
type ShoppingListService struct {
	createShoppingListUseCase usecase.ICreateShoppingListUseCase
}

func (s ShoppingListService) CreateNewShoppingList(ctx context.Context, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error) {
	return s.createShoppingListUseCase.CreateNewShoppingList(ctx, req)
}

func NewShoppingListService(createShoppingListUseCase usecase.ICreateShoppingListUseCase) IShoppingListService {
	return &ShoppingListService{
		createShoppingListUseCase: createShoppingListUseCase,
	}
}
