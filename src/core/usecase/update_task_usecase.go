package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IUpdateTaskUsecase interface {
	UpdateTaskByID(ctx context.Context, userID, shoppingListID, taskID int64, updateTaskDto *request.UpdateTaskDto) (*entity.ShoppingTaskEntity, error)
}

type UpdateTaskUsecase struct {
	shoppingTaskPort           port.IShoppingTaskPort
	getShoppingListUsecase     IGetShoppingListUsecase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (u UpdateTaskUsecase) UpdateTaskByID(ctx context.Context, userID, shoppingListID, taskID int64, updateTaskDto *request.UpdateTaskDto) (*entity.ShoppingTaskEntity, error) {
	_, err := u.getShoppingListUsecase.GetShoppingListByID(ctx, userID, shoppingListID)
	if err != nil {
		log.Error(ctx, "UpdateTaskByID: GetShoppingListByID error", err)
		return nil, err
	}
	task, err := u.shoppingTaskPort.GetTaskByID(ctx, taskID)
	if err != nil {
		log.Error(ctx, "UpdateTaskByID: GetTaskByID error", err)
		return nil, err
	}
	if updateTaskDto.FoodName != nil {
		task.FoodName = *updateTaskDto.FoodName
	}
	if updateTaskDto.Quantity != nil {
		task.Quantity = *updateTaskDto.Quantity
	}
	if updateTaskDto.Status != nil {
		task.Status = *updateTaskDto.Status
	}
	tx := u.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error(ctx, "UpdateTaskByID: Rollback error", errRollback)
		} else {
			log.Info(ctx, "UpdateTaskByID: Rollback successfully")
		}
	}()
	task, err = u.shoppingTaskPort.UpdateTaskByID(ctx, tx, task)
	if err != nil {
		log.Error(ctx, "UpdateTaskByID: UpdateTaskByID error", err)
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		log.Error(ctx, "UpdateTaskByID: Commit error", err)
		return nil, err
	}
	return task, nil
}

func NewUpdateTaskUsecase(shoppingTaskPort port.IShoppingTaskPort, getShoppingListUsecase IGetShoppingListUsecase, databaseTransactionUsecase IDatabaseTransactionUsecase) IUpdateTaskUsecase {
	return &UpdateTaskUsecase{
		shoppingTaskPort:           shoppingTaskPort,
		getShoppingListUsecase:     getShoppingListUsecase,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
