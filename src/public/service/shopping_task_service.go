package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IShoppingTaskService interface {
	CreateNewShoppingTask(ctx context.Context, userID, shoppingListId int64, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error)
	GetShoppingTasksByShoppingListID(ctx context.Context, userID, shoppingListID int64) ([]*entity.ShoppingTaskEntity, error)
	DeleteTaskByID(ctx context.Context, userID, shoppingListID, taskID int64) error
}
type ShoppingTaskService struct {
	createShoppingTaskUsecase usecase.ICreateShoppingTaskUsecase
	getTaskUsecase            usecase.IGetTaskUsecase
	deleteTaskUsecase         usecase.IDeleteTaskUsecase
}

func (s ShoppingTaskService) DeleteTaskByID(ctx context.Context, userID, shoppingListID, taskID int64) error {
	return s.deleteTaskUsecase.DeleteTaskByID(ctx, userID, shoppingListID, taskID)
}

func (s ShoppingTaskService) GetShoppingTasksByShoppingListID(ctx context.Context, userID, shoppingListID int64) ([]*entity.ShoppingTaskEntity, error) {
	return s.getTaskUsecase.GetTaskByShoppingListID(ctx, userID, shoppingListID)
}

func (s ShoppingTaskService) CreateNewShoppingTask(ctx context.Context, userID, shoppingListId int64, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error) {
	return s.createShoppingTaskUsecase.CreateNewShoppingTask(ctx, userID, shoppingListId, shoppingTasks)
}

func NewShoppingTaskService(createShoppingTaskUsecase usecase.ICreateShoppingTaskUsecase, getTaskUsecase usecase.IGetTaskUsecase, deleteTaskUsecase usecase.IDeleteTaskUsecase) IShoppingTaskService {
	return &ShoppingTaskService{createShoppingTaskUsecase, getTaskUsecase, deleteTaskUsecase}
}
