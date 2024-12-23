package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IDeleteFridgeItemUsecase interface {
	DeleteItem(ctx context.Context, userID, itemID int64) error
}
type DeleteFridgeItemUsecase struct {
	fridgeItemPort             port.IFridgeItemPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (d DeleteFridgeItemUsecase) DeleteItem(ctx context.Context, userID, itemID int64) error {
	item, err := d.fridgeItemPort.GetItemByIDAndCreatedBy(ctx, itemID, userID)
	if err != nil {
		log.Error(ctx, "Get fridge item failed: ", err)
		return err
	}
	tx := d.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error(ctx, "Rollback failed: ", errRollback)
		} else {
			log.Info(ctx, "Rollback successfully")
		}
	}()
	err = d.fridgeItemPort.DeleteItem(ctx, tx, item.ID)
	if err != nil {
		log.Error(ctx, "Delete fridge item failed: ", err)
		return err
	}
	if err = tx.Commit().Error; err != nil {
		log.Error(ctx, "Commit failed: ", err)
		return err
	}
	return nil
}

func NewDeleteFridgeItemUsecase(fridgeItemPort port.IFridgeItemPort) IDeleteFridgeItemUsecase {
	return &DeleteFridgeItemUsecase{
		fridgeItemPort: fridgeItemPort,
	}
}
