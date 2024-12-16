package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IShoppingListService interface {
	CreateNewShoppingList(ctx context.Context, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error)
	UpdateShoppingList(ctx context.Context, userID int64, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error)
	DeleteShoppingListByID(ctx context.Context, userID, shoppingListId int64) error
}
type ShoppingListService struct {
	createShoppingListUseCase usecase.ICreateShoppingListUseCase
	updateShoppingListUseCase usecase.IUpdateShoppingListUseCase
	deleteShoppingListUseCase usecase.IDeleteShoppingListUseCase
}

func (s ShoppingListService) DeleteShoppingListByID(ctx context.Context, userID int64, id int64) error {
	return s.deleteShoppingListUseCase.DeleteShoppingListByID(ctx, userID, id)
}

func (s ShoppingListService) UpdateShoppingList(ctx context.Context, userID int64, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error) {
	return s.updateShoppingListUseCase.UpdateShoppingListByID(ctx, userID, req)
}

func (s ShoppingListService) CreateNewShoppingList(ctx context.Context, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error) {
	return s.createShoppingListUseCase.CreateNewShoppingList(ctx, req)
}

func NewShoppingListService(
	createShoppingListUseCase usecase.ICreateShoppingListUseCase,
	updateShoppingListUseCase usecase.IUpdateShoppingListUseCase,
	deleteShoppingListUseCase usecase.IDeleteShoppingListUseCase,
) IShoppingListService {
	return &ShoppingListService{
		createShoppingListUseCase: createShoppingListUseCase,
		updateShoppingListUseCase: updateShoppingListUseCase,
		deleteShoppingListUseCase: deleteShoppingListUseCase,
	}
}
