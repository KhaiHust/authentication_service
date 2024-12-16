package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IDeleteTaskUsecase interface {
	DeleteTaskByID(ctx context.Context, userID, shoppingListID, taskID int64) error
}
type DeleteTaskUsecase struct {
	shoppingTaskPort           port.IShoppingTaskPort
	getShoppingListUsecase     IGetShoppingListUsecase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (d DeleteTaskUsecase) DeleteTaskByID(ctx context.Context, userID, shoppingListID, taskID int64) error {
	_, err := d.getShoppingListUsecase.GetShoppingListByID(ctx, userID, shoppingListID)
	if err != nil {
		log.Error(ctx, "Get shopping list by id error: ", err)
		return err
	}
	tx := d.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := d.databaseTransactionUsecase.Rollback(tx); errRollback != nil {
			log.Error(ctx, "Rollback transaction error: ", errRollback)
		} else {
			log.Info(ctx, "Rollback transaction success")
		}
	}()
	err = d.shoppingTaskPort.DeleteTaskByID(ctx, tx, taskID)
	if err != nil {
		log.Error(ctx, "Delete task by id error: ", err)
		return err
	}
	errCommit := d.databaseTransactionUsecase.Commit(tx)
	if errCommit != nil {
		log.Error(ctx, "Commit transaction error: ", errCommit)
		return errCommit
	}
	return nil
}

func NewDeleteTaskUsecase(shoppingTaskPort port.IShoppingTaskPort, getShoppingListUsecase IGetShoppingListUsecase, databaseTransactionUsecase IDatabaseTransactionUsecase) IDeleteTaskUsecase {
	return &DeleteTaskUsecase{shoppingTaskPort, getShoppingListUsecase, databaseTransactionUsecase}
}
